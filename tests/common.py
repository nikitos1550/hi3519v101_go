import os
import sys
import subprocess
import logging


TESTS_DIR = os.path.dirname(__file__)
PROJECT_DIR = os.path.abspath(os.path.join(TESTS_DIR, ".."))
BR_HISICAM_TESTENV_DIR = os.path.abspath(os.path.join(PROJECT_DIR, "br-hisicam/tests"))


if not os.path.isdir(BR_HISICAM_TESTENV_DIR):
    raise SystemExit(
        "br-hisicam/tests directory is absent. Make sure 'br-hisicam' submodule is initialized"
    )

sys.path.insert(0, BR_HISICAM_TESTENV_DIR)


from testenv import DEVICE_LIST, br_hisicam, hiburn



# -------------------------------------------------------------------------------------------------
class Make:
    def __init__(self):
        pass

    def info(self, board):
        out = subprocess.check_output([
            "make", "NO_USER_MAKEFILE=Y", f"BOARD={board}", "APP_TARGET=cam", "info"
        ], cwd=PROJECT_DIR)

        info = {}
        for line in out.decode("utf-8").split("\n"):
            line = line.strip()
            if not line:
                continue
            try:
                key, value = line.split("=")
                info[key] = value
            except ValueError as err:
                logging.error(f"Couldn't unpack {line}: {err}")
        return info

    def build_app(self, board):
        subprocess.check_call([
            "make", "NO_USER_MAKEFILE=Y", f"BOARD={board}", "APP_TARGET=cam", "build-app"
        ], cwd=PROJECT_DIR)


make = Make()