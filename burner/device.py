#!/usr/bin/env python3
import serial
import sys


class Device:
    def _log_data(self, prefix, data):
        if self.logger is None:
            return
        if isinstance(data, bytes):
            data = data.decode("ascii")
        for r in (("\n", "\\n"), ("\r", "\\r")):
            data = data.replace(r[0], r[1])
        self.logger.info(prefix + data)

    def log_out(self, data):
        self._log_data("-> ", data)
    
    def log_in(self, data):
        self._log_data("<- ", data)

    def dlog(self, message, *args, **kwargs):
        if self.logger is not None:
            self.logger.debug(str(self) + ": " + message.format(*args, **kwargs))

    def __str__(self):
        return self.serial.port

    def __del__(self):
        self.close()

    def __init__(self, port, baudrate, read_timeout=None, logger=None):
        self.logger = logger
        self.serial = serial.Serial(
            port=port,
            baudrate=baudrate,
            timeout=read_timeout
        )
        self.dlog("serial connection made")
    
    def close(self):
        if hasattr(self, "serial"):
            self.serial.close()
            self.dlog("serial connection closed")

    def clear_input_buff(self):
        self.serial.reset_input_buffer()
        self.dlog("clear input buffer")

    def write_data(self, data):
        self.serial.write(data)
        self.log_out(data)

    def read_line(self):
        line = self.serial.readline()
        if line:
            self.log_in(line)
        return line


def read(dev, args):
    dev.dlog("Read from {}...", dev)
    while True:
        line = dev.read_line()
        try:
            line = line.decode("utf-8")
        except:
            pass
        sys.stdout.write(line)
        sys.stdout.flush()


def write(dev, args):
    if args.wait_for is not None:
        dev.dlog("Wait for '{}' from {}...", args.wait_for, dev)
        while args.wait_for not in dev.read_line().decode("ascii"):
            pass
    dev.dlog("Write '{}' to {}...", args.data, dev)
    dev.write_data(args.data.encode("ascii") + b"\n")


def console(dev, args):
    import signal, termios, tty, asyncio

    dev.serial.timeout = 0.05

    # prepare TTY
    term_fd = sys.stdin.fileno()
    orig_term_attrs = termios.tcgetattr(term_fd)
    tty_attrs = orig_term_attrs[:]
    tty_attrs[tty.LFLAG] &= ~(termios.ECHO | termios.ICANON)
    termios.tcsetattr(term_fd, termios.TCSADRAIN, tty_attrs)

    # handle Ctrl+C
    running = True
    def signal_handler(signal, frame):
        nonlocal running
        running = False
    signal.signal(signal.SIGINT, signal_handler)

    def on_stdin():
        try:
            s = sys.stdin.read(1).encode("ascii")
            dev.write_data(s)
        except: pass

    async def read_from_dev():
        dev.write_data(b"\n")
        while running:
            line = dev.read_line()
            if line:
                sys.stdout.write(line.decode("ascii"))
            else:
                sys.stdout.flush()
                await asyncio.sleep(0.1)
        
    loop = asyncio.get_event_loop()
    loop.add_reader(sys.stdin, on_stdin)
    loop.run_until_complete(read_from_dev())

    # restore terminal
    termios.tcsetattr(term_fd, termios.TCSAFLUSH, orig_term_attrs)


def main():
    import argparse
    import utils
    
    parser = argparse.ArgumentParser(description="Simple interaction via serial port")
    parser.add_argument("--port", required=True, help="Device that represents serial connection")
    parser.add_argument("--br", type=int, help="Serial baud rate", metavar="BAUDRATE", default=115200)
    parser.add_argument("--verbose", action="store_true", help="Enable logging", default=False)
    parser.set_defaults(act=console)

    action_parsers = parser.add_subparsers(title="Action")

    # read
    read_parser = action_parsers.add_parser("read", help="Read data from serial port and print it")
    read_parser.set_defaults(act=read)

    # write
    write_parser = action_parsers.add_parser("write", help="Write data into serial port")
    write_parser.add_argument("--wait-for", metavar="LINE", help="Write data when LINE has been received", required=False)
    write_parser.add_argument("--data", help="Data to be written into serial port", required=False)
    write_parser.set_defaults(act=write)

    # console
    console_parser = action_parsers.add_parser("console", help="Simple interactive console")
    console_parser.set_defaults(act=console)

    args = parser.parse_args()

    logger = utils.get_device_logger(args.port) if args.verbose else None
    device = Device(port=args.port, baudrate=args.br, logger=logger)
    args.act(device, args)


if __name__ == "__main__":
    main()
