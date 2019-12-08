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

    req = " ".join(command).encode("ascii") + b"\n"

    writer.write(req)
    try:
        resp = (await reader.readuntil(b"#"))[:-2]
    except asyncio.streams.IncompleteReadError as err:
        logging.warning("Request '{}' seems failed".format(req))
        resp = err.partial

    writer.close()
    return json.loads(resp.decode("ascii"))


def print_response(resp):
    status = resp["status"]
    common.print_success(status) if (status == common.STATUS_OK) else common.print_failure("OK")
    print(" >> {}".format(resp["message"]))


def handle_response(resp):
    print_response(resp)
    if resp["exec"]:
        args = resp["exec"].split(" ")
        if args[0] == "telnet":
            subprocess.run(args, shell=False)
            common.print_success("Telnet session completed\n")


def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--port", "-p", type=int, default=43500, help="TCP port of the daemon")
    parser.add_argument("command", nargs="+", type=str, help="Command and arguments")

    args = parser.parse_args()

    eloop = asyncio.get_event_loop()
    response = eloop.run_until_complete(request(command=args.command, port=args.port))
    handle_response(response)


main()
