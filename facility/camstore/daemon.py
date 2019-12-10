#!/bin/env python3

import os
import sys
import logging


STOP_SIGNALS = ("SIGINT", "SIGTERM")  # this signals will stop server
LOG_FILE = os.path.join(os.path.dirname(__file__), "log")  # default log file
PID_FILE = os.path.join(os.path.dirname(__file__), ".camstore.pid")  # default PID file


def run(port, config):
    import asyncio
    import signal

    from lib.server import Server
    from lib import info, ser2net_controller, devices

    logging.info("--- STARTED")

    srv = Server()

    eloop = asyncio.get_event_loop()
    for signame in STOP_SIGNALS:
        eloop.add_signal_handler(getattr(signal, signame), srv.stop)

    try:
        eloop.run_until_complete(srv.run(port=port))
    except Exception as err:
        logging.fatal("Server not started: {}".format(err))

    # needed by some asynchronous reasons...
    for signame in STOP_SIGNALS:
        eloop.remove_signal_handler(getattr(signal, signame))

    logging.info("--- FINISHED")


def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--config", "-c", type=str, help="Configuration file")
    parser.add_argument("--port", "-p", type=int, default=43500, help="Port to listen on")
    parser.add_argument("--detach", "-d", action="store_true", help="Fork process")
    parser.add_argument("--pidf", type=str, help="File to write PID into")
    parser.add_argument("--logf", type=str, help="File to write logs into")

    args = parser.parse_args()

    if args.detach:
        pid = os.fork()
        if pid:
            print("PID: {}".format(pid))
            exit(0)
        if not args.logf:
            args.logf = LOG_FILE
        if not args.pidf:
            args.pidf = PID_FILE

    logging.basicConfig(level=logging.DEBUG, filename=args.logf, filemode="w")
    if args.pidf:
        with open(args.pidf, "w") as f:
            f.write(str(os.getpid()))

    run(args.port, args.config)


main()
