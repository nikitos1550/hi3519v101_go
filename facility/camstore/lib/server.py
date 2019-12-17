import asyncio
import inspect
import logging
import functools
from .common import success, failure, InvalidArgument


COMMANDS = {}
ROUTINES = {}


def register(func):
    """ Decorator: function will be available as a command
    """

    cmd_name = func.__name__.lower()
    if cmd_name in COMMANDS:
        raise Exception("Command {} already registered".format(cmd_name))
    
    @functools.wraps(func)
    async def wrap(context, *args, **kwargs):
        # take needed arguments from context
        sign = inspect.signature(func)
        for k, v in context.items():
            if k in sign.parameters:
                kwargs[k] = v
        try:
            if inspect.iscoroutinefunction(func):
                res = await func(*args, **kwargs)
            else:
                res = func(*args, **kwargs)
            return res
        except InvalidArgument as err:
            return failure("Invalid argument: {}".format(err))

    COMMANDS[cmd_name] = wrap
    return wrap


def routine(func):
    """ Decorator: function will be scheduled as a task
    """

    name = func.__name__

    async def wrap():
        while True:
            try:
                logging.info(f"Start '{name}' routine...")
                await func()
            except asyncio.CancelledError:
                logging.info(f"Routine '{name}' cancelled")
                return
            except Exception as err:
                logging.error(f"Routine '{name}' failed with '{err.__class__.__name__}': {err}")
                logging.info(f"Sleep 3 seconds before restart '{name}' routine")
                await asyncio.sleep(3)

    ROUTINES[func.__name__] = wrap
    return wrap


# -------------------------------------------------------------------------------------------------
@register
def help(cmd_name=None):
    """Print available commands
    """
    def command_list():
        for name, func in COMMANDS.items():
            if func.__doc__:
                title = func.__doc__.split("\n")[0].strip()
            else:
                title = "untitled"
            yield "  {} - {}".format(name, title)

    if cmd_name:
        cmd = COMMANDS.get(cmd_name)
        if cmd is None:
            return failure("command {} does not exist".format(cmd_name))
        doc = cmd.__doc__.strip() or "No help for command {}".format(cmd_name)
    else:
        doc = "Available commands:\n" + "\n".join(l for l in command_list())
    return success(message=doc)


# -------------------------------------------------------------------------------------------------
class Connection:
    @property
    def peer(self):
        return self.writer.transport.get_extra_info("peername")

    def __init__(self, reader, writer):
        self.reader = reader
        self.writer = writer
        self.context = {}

    def write(self, data):
        self.writer.write(str(data).encode("ascii") + b"\n")

    def set_ready(self):
        self.writer.write(b"#")

    async def process_command(self, line):
        logging.debug(f"Received command: {line}")

        args = line.split(" ")
        cmd_name = args[0].lower()

        if cmd_name in ("exit", "quit"):
            self.write(success("bye"))
            return

        if cmd_name == "set_user":
            self.context["user"] = args[1]
            logging.info("Set connection user '{}'".format(self.context["user"]))
            return success("User '{}' set".format(self.context["user"]))

        cmd = COMMANDS.get(cmd_name)
        if cmd is None:
            return failure(f"unknown command '{cmd_name}'")

        res = cmd(self.context, *args[1:])
        if inspect.iscoroutine(res):
            res = await res
        if res is None:
            res = success("OK")
        logging.debug(f"Response: {res}")
        return res

    async def run(self):
        logging.info("Connection from {}:{} opened".format(*self.peer))
        try:
            while True:
                self.set_ready()
                line = (await self.reader.readline()).decode("ascii").strip()
                if not line:
                    if self.reader.at_eof():
                        return
                    continue

                response = await self.process_command(line)
                if response is not None:
                    self.write(response)
                else:
                    break
        except asyncio.CancelledError as err:
            self.write(failure("Server stopped"))
        except Exception as err:
            logging.exception("", exc_info=err)
            self.write(failure(f"Internal error occurred: {err}"))
        finally:
            logging.info("Connection from {}:{} closed".format(*self.peer))
            self.writer.close()



# -------------------------------------------------------------------------------------------------
class Server:
    def __init__(self):
        self._srv = None
        self._tasks = set()

    def stop(self, *_):
        if self._srv is not None:
            self._srv.close()
            logging.info("Server has closed")

        for t in self._tasks:
            if not t.done():
                t.cancel()

    async def run(self, port):
        try:
            self._srv = await asyncio.start_server(self._on_connect, host="localhost", port=port)
            logging.info("Server is listening on localhost:{}".format(port))

            self._start_routines()
            logging.info("Routines are started")

            await self._srv.wait_closed()
            logging.info("Server has stopped")

            await asyncio.gather(*self._tasks, return_exceptions=True)
            logging.info("All server's tasks completed")
        except:
            self.stop()
            raise

    async def _on_connect(self, reader, writer):
        task = asyncio.Task.current_task()
        self._tasks.add(task)
        try:
            await Connection(reader, writer).run()
        finally:
            self._tasks.remove(task)

    def _start_routines(self):
        for r in ROUTINES.values():
            self._tasks.add(asyncio.ensure_future(r()))
