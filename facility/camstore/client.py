#!/bin/env python3

import asyncio
import logging
import subprocess
import os, sys
import json

if __package__:
    from .lib import common
else:
    from lib import common


ALLOWED_EXEC_COMMANDS = ("telnet")


class FailedRequest(Exception):
    def __init__(self, response):
        self.response = response


class Connection:
    @classmethod
    async def connect(cls, port):
        reader, writer = await asyncio.open_connection(host="localhost", port=port)
        await reader.readuntil(b"#")
        return cls(reader, writer)

    def __enter__(self):
        return self

    def __exit__(self, *arg, **kwargs):
        self.close()

    def __init__(self, reader, writer):
        self.reader = reader
        self.writer = writer

    async def _make_request(self, request):
        self.writer.write(request.encode("ascii") + b"\n")
        try:
            response = await self.reader.readuntil(b"#")
            return response[:-2].decode("ascii")
        except asyncio.IncompleteReadError as err:
            return err.partial.decode("ascii").strip()

    async def request(self, req):
        resp = json.loads(await self._make_request(req))
        if resp["status"] != common.STATUS_OK:
            raise FailedRequest(resp)
        return resp
    
    def close(self):
        self.writer.close()

# -------------------------------------------------------------------------------------------------

async def request(command, port):
    conn = await Connection.connect(port)
    with conn:
        await conn.request("set_user {}".format(os.getlogin()))
        return await conn.request(" ".join(command))


def handle_response(resp, execute=True):
    common.print_response(resp)
    if resp.get("status") != common.STATUS_OK:
        return -1
    
    if not resp.get("exec"):
        return 0

    args = resp["exec"].split(" ")

    if args[0] not in ALLOWED_EXEC_COMMANDS:
        common.print_failure(f"Command '{args[0]}' is not allowed")
        return -1

    if not execute:
        print(resp["exec"])
        return 0

    subprocess.run(args, shell=False)
    common.print_success(f"{args[0]} session completed\n")

    return 0

# -------------------------------------------------------------------------------------------------

def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--port", "-p", type=int, default=43500, help="TCP port of the daemon")
    parser.add_argument("--no-exec", action="store_true", help="Do not execute (telnet)")
    parser.add_argument("command", nargs="+", type=str, help="Command and arguments")

    args = parser.parse_args()

    logging.basicConfig(level=logging.WARNING)

    eloop = asyncio.get_event_loop()
    try:
        response = eloop.run_until_complete(request(command=args.command, port=args.port))
        rc = handle_response(response, execute=(not args.no_exec))
        exit(rc)
    except FailedRequest as err:
        common.print_response(err.response)
        exit(-1)
    except Exception as err:
        logging.error(err)
        exit(-2)


if __name__ == "__main__":
    main()
