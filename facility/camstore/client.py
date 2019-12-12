#!/bin/env python3

import asyncio
import logging
import subprocess
import os, sys
import json
from lib import common


logging.basicConfig(level=logging.WARNING)


async def request(command, port):
    try:
        reader, writer = await asyncio.open_connection(host="localhost", port=port)
        await reader.readuntil(b"#")
    except Exception as err:
        logging.error("Could not connect to localhost:{}".format(port))
        return

    writer.write("set_user {}".format(os.getlogin()).encode("ascii") + b"\n")
    await reader.readuntil(b"#")

    req = " ".join(command).encode("ascii") + b"\n"
    writer.write(req)
    try:
        resp = (await reader.readuntil(b"#"))[:-2]
    except asyncio.streams.IncompleteReadError as err:
        logging.warning("Request '{}' seems failed".format(req))
        resp = err.partial

    writer.close()
    return json.loads(resp.decode("ascii"))


def handle_response(resp, execute=True):
    common.print_response(resp)
    if resp.get("status") != common.STATUS_OK:
        return -1

    if resp["exec"]:
        if not execute:
            print(resp["exec"])
        else:
            args = resp["exec"].split(" ")
            if args[0] == "telnet":
                subprocess.run(args, shell=False)
                common.print_success("Telnet session completed\n")

    return 0

# -------------------------------------------------------------------------------------------------

def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--port", "-p", type=int, default=43500, help="TCP port of the daemon")
    parser.add_argument("--no-exec", action="store_true", help="Do not execute (telnet)")
    parser.add_argument("command", nargs="+", type=str, help="Command and arguments")

    args = parser.parse_args()

    eloop = asyncio.get_event_loop()
    response = eloop.run_until_complete(request(command=args.command, port=args.port))
    if response is None:
        exit(-1)

    rc = handle_response(response, execute=(not args.no_exec))
    exit(rc)


main()
