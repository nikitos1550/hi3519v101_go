#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Defaults:
    dev_baudrate = 115200


# =====================================================================================================================
class MacAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("mac", help="Change MAC address")
        parser.add_argument("--val", help="MAC address to be setup (random by default)")
        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(args):
        print("Wanna change {} MAC on {} you bastard?".format(args.port, args.val))


# =====================================================================================================================
class LoadAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("load", help="Load image onto device and boot it (without burning)")
        parser.add_argument("--uimage", help="Kernel's uImage file")
        parser.add_argument("--rootfs", help="RootFS file")

        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(args):
        print("Wanna boot {} from {} you bastard?".format(args.port, args.uimage))


# =====================================================================================================================
def main():
    import argparse

    parser = argparse.ArgumentParser(description="Interact with devices via serial port")
    parser.add_argument("--port", required=True, help="Serial port dev name")
    parser.add_argument("--baudrate", type=int, default=Defaults.dev_baudrate,
        help="Serial port baudrate (default: {})".format(Defaults.dev_baudrate))
    parser.set_defaults(action=lambda args: parser.print_help())

    action_parsers = parser.add_subparsers(title="Action")
    MacAction.register(action_parsers)
    LoadAction.register(action_parsers)
    
    args = parser.parse_args()
    if args.action is not None:
        args.action(args)


if __name__ == "__main__":
    main()


