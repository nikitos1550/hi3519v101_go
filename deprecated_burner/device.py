#!/usr/bin/env python3
import serial
import sys
from telnetlib import Telnet


# -------------------------------------------------------------------------------------------------
# Connect to a device's serial port using pyserial
class SerialConn:
    def __str__(self):
        return f"serial({self.conn.port})"

    def __init__(self, port, baudrate, read_timeout=None):
        self.conn = serial.Serial(
            port=port,
            baudrate=baudrate,
            timeout=read_timeout
        )

    def set_read_timeout(self, timeout):
        self.conn.timeout = timeout

    def close(self):
        self.conn.close()

    def reset_input_buffer(self):
        self.conn.reset_input_buffer()

    def write_data(self, data):
        self.conn.write(data)

    def read_line(self):
        return self.conn.readline()


# -------------------------------------------------------------------------------------------------
# Connect to a device's telnet port
class TelnetConn:
    def __str__(self):
        return f"telnet({self.host}:{self.port})"

    def __init__(self, port, host="localhost", read_timeout=None):
        self.host = host
        self.port = port
        self.conn = Telnet(host=self.host, port=self.port)
        self._read_timeout = read_timeout

    def set_read_timeout(self, timeout):
        self._read_timeout = timeout

    def close(self):
        self.conn.close()

    def reset_input_buffer(self):
        self.conn.read_very_eager()

    def write_data(self, data):
        self.conn.write(data)

    def read_line(self):
        return self.conn.read_until(b"\n", timeout=self._read_timeout)


# -------------------------------------------------------------------------------------------------
class Device:
    def _log_data(self, prefix, data):
        if self.logger is None:
            return
        try:
            if isinstance(data, bytes):
                data = data.decode("ascii")
            for r in (("\n", "\\n"), ("\r", "\\r")):
                data = data.replace(r[0], r[1])
            self.logger.debug(prefix + data)
        except:
            pass

    def log_out(self, data):
        self._log_data("-> ", data)
    
    def log_in(self, data):
        self._log_data("<- ", data)

    def dlog(self, message, *args, **kwargs):
        if self.logger is not None:
            self.logger.info(f"{self}: " + message.format(*args, **kwargs))

    def __str__(self):
        return self.conn.__str__()

    def __del__(self):
        self.close()

    def __init__(self, conn, logger=None):
        self.logger = logger
        self.conn = conn
        self.dlog("connection made")
    
    def close(self):
        if hasattr(self, "conn"):
            self.conn.close()
            self.dlog("connection closed")

    def clear_input_buff(self):
        self.conn.reset_input_buffer()
        self.dlog("clear input buffer")

    def write_data(self, data):
        self.conn.write_data(data)
        self.log_out(data)

    def read_line(self):
        line = self.conn.read_line()
        if line:
            self.log_in(line)
        return line


# -------------------------------------------------------------------------------------------------
def read(dev, args):
    dev.dlog("Read from {}...", dev)
    while True:
        line = dev.read_line()
        try:
            line = line.decode("utf-8")
        except:
            pass
        sys.stdout.write(str(line))
        sys.stdout.flush()


def write(dev, args):
    if args.wait_for is not None:
        dev.dlog("Wait for '{}' from {}...", args.wait_for, dev)
        while args.wait_for not in dev.read_line().decode("ascii"):
            pass
    dev.dlog("Write '{}' to {}...", args.data, dev)
    dev.write_data(args.data.encode("ascii") + b"\n")


def main():
    import argparse
    import utils
    
    parser = argparse.ArgumentParser(description="Simple interaction via serial port")
    parser.add_argument("--port", required=True, help="Device that represents serial connection")
    parser.add_argument("--telnet", action="store_true", help="COnnect via Telnet")
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

    args = parser.parse_args()

    logger = utils.get_device_logger(args.port) if args.verbose else None
    if args.telnet:
        conn = TelnetConn(port=args.port)
    else:
        conn = SerialConn(port=args.port, baudrate=args.br)
    device = Device(conn=conn, logger=logger)
    args.act(device, args)


if __name__ == "__main__":
    main()
