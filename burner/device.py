#!/usr/bin/env python
import serial
from utils import CLOG_INCOMING, CLOG_OUTGOING, DLOG_INFO


class Device:
    def __init__(self, port, baudrate, timeout=0.1):
        self._serial_port = serial.Serial(
            port=port,
            baudrate=baudrate,
            timeout=timeout
        )
        DLOG_INFO("Serial connection for {} made".format(port))
    
    def close():
        self._serial_port.close()

    def write_data(self, data):
        self._serial_port.write(data)
        CLOG_OUTGOING(data)

    def write_cmd(self, cmd):
        data = (cmd.replace(";", "\;") + "\n").encode("ascii")
        self.write_data(data)

    def write_ctrlc(self):
        """ Send Ctrl+C to device
        """
        self.write_data(CTRL_C)

    def wait_prompt(self, clear=True):
        """If 'clear' then we wait for a line contains only prompt, otherwise any
        line with prompt is fine.
        """
        while True:
            line = self.read_line()
            if clear:
                if (line in PROMPTS):
                    break
            elif line_has_prompt(line):
                break
            
    def read_line(self):
        line = self._serial_port.readline().strip()
        if line:
            CLOG_INCOMING(line)
        return line


def read(dev, args):
    DLOG_INFO("Read from {}...".format(dev))
    while True:
        dev.read_line()


def write(dev, args):
    if args.wait_for is not None:
        DLOG_INFO("Wait for '{}' from {}...".format(args.wait_for, dev))
        while dev.read_line() != args.wait_for: pass
    DLOG_INFO("Write '{}' to {}...".format(args.data, dev))
    dev.write_data(args.data + "\n")


def main():
    import argparse
    
    parser = argparse.ArgumentParser(description="Simple interaction via serial port")
    parser.add_argument("--port", required=True, help="Device that represents serial connection")
    parser.add_argument("--br", required=True, type=int, help="Serial baud rate", metavar="BAUDRATE", default=115200)

    action_parsers = parser.add_subparsers(title="Action")

    # read action
    read_parser = action_parsers.add_parser("read", help="Read data from serial port and print it")
    read_parser.set_defaults(act=read)

    # write action
    write_parser = action_parsers.add_parser("write", help="Write data into serial port")
    write_parser.add_argument("--wait-for", help="Write data when this has been received", required=False)
    write_parser.add_argument("--data", help="Data to be written into serial port", required=False)

    write_parser.set_defaults(act=write)

    args = parser.parse_args()

    device = Device(port=args.port, baudrate=args.br)
    args.act(device, args)


if __name__ == "__main__":
    main()
