import json
import sys
import os


class InvalidArgument(Exception):
    pass


STATUS_OK = "OK"
STATUS_FAILED = "FAILED"


class Response:
    def __init__(self, status, message=None, execute=None):
        self.res = dict(status=status, message=message, exec=execute)

    def __str__(self):
        return json.dumps(self.res)


def success(message=None, execute=None):
    """ Create and return success response
    """
    return Response(status=STATUS_OK, message=message, execute=execute)


def failure(message=None):
    """ Create and return failure response
    """
    return Response(status=STATUS_FAILED, message=message)


if os.isatty(sys.stdout.fileno()):
    COLOR_RED       = "\033[0;31m"
    COLOR_GREEN     = "\033[0;32m"
    COLOR_DEFAULT   = "\033[0m"

    def _color_print(color, message):
        sys.stdout.write(color)
        sys.stdout.write(message)
        sys.stdout.write(COLOR_DEFAULT)

    def print_failure(msg):
        _color_print(COLOR_RED, msg)

    def print_success(msg):
        _color_print(COLOR_GREEN, msg)

else:
    def print_failure(msg):
        sys.stdout.write(msg)

    def print_success(msg):
        sys.stdout.write(msg)


def print_response(response):
    status = response["status"]
    if status == STATUS_OK:
        print_success(status)
    else:
        print_failure(status)
    if "message" in response:
        print(" >> {}".format(response["message"]))
