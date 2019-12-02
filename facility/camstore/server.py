import asyncio
import inspect
import logging
import functools
import common


COMMANDS = {}
ROUTINES = {}


def register(func):
    """ Decorator: function will be available as a command
    """

    cmd_name = func.__name__.lower()
    if cmd_name in COMMANDS:
        raise Exception("Command {} already registered".format(cmd_name))
    
    @functools.wraps(func)
    async def wrap(*args, **kwargs):
        try:
            if inspect.iscoroutinefunction(func):
                res = await func(*args, **kwargs)
            else:
                res = func(*args, **kwargs)
            return res
        except common.InvalidArgument as err:
            return "failed: invalid argument - {}".format(err)

    COMMANDS[cmd_name] = wrap
    return wrap


def routine(func):
    """ Decorator: function will be scheduled as a task
    """
    async def wrap():
        try:
            await func()
        except Exception as err:
            logging.error("Routine {} failed with {}: {}".format(func.__name__, err.__class__.__name__, err))
        
    asyncio.CancelledError
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
        if cmd:
            return cmd.__doc__ or "No help for the command"
        else:
            return "failed: command {} does not exist".format(cmd_name)
    else:
        return "Available commands:\n" + "\n".join(l for l in command_list())


# -------------------------------------------------------------------------------------------------
class Connection:
    @property
    def peer(self):
        return self.writer.transport.get_extra_info("peername")

    def __init__(self, reader, writer):
        self.reader = reader
        self.writer = writer

    def write(self, data):
        self.writer.write(data.encode("ascii") + b"\n")

    def set_ready(self):
        self.writer.write(b"#")

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

                args = line.split(" ")
                cmd_name = args[0].lower()

                if cmd_name in ("exit", "quit"):
                    self.write("ok: bye")
                    break

                cmd = COMMANDS.get(cmd_name)
                if cmd is not None:
                    if inspect.iscoroutinefunction(cmd):
                        res = await cmd(*args[1:])
                    else:
                        res = cmd(*args[1:])
                    self.write(res)
                else:
                    self.write("failed: unknown command {}".format(cmd_name))

        except asyncio.CancelledError as err:
            self.write("server stopped")
        except Exception as err:
            logging.exception("", exc_info=err)
            self.write("failed: internal error occurred {}".format(err))
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
