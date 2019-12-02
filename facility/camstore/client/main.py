#!/bin/env python3

import asyncio
import logging
import subprocess


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
    return resp.decode("ascii")


def handle_response(resp: str):
    if resp.startswith("ok/exec:"):
        subprocess.run(resp.split(":")[1].strip(), shell=True)
    else:
        print(resp)



def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--port", "-p", type=int, required=True, help="TCP port of the daemon")
    parser.add_argument("command", nargs="+", type=str, help="Command and arguments")

    args = parser.parse_args()

    eloop = asyncio.get_event_loop()
    response = eloop.run_until_complete(request(command=args.command, port=args.port))
    handle_response(response)


main()
