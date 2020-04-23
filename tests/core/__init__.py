import os
import sys


# -------------------------------------------------------------------------------------------------
TEST_CORE_DIR = os.path.dirname(__file__)
PROJECT_DIR = os.path.abspath(os.path.join(TEST_CORE_DIR, "../.."))
BR_HISICAM_TESTS_DIR = os.path.abspath(os.path.join(PROJECT_DIR, "br-hisicam/tests"))


if not os.path.isdir(BR_HISICAM_TESTS_DIR):
    raise SystemExit(
        "br-hisicam/tests directory is absent. Make sure 'br-hisicam' submodule is initialized"
    )

sys.path.insert(0, BR_HISICAM_TESTS_DIR)


from testenv import DEVICE_LIST, br_hisicam, hiburn
from .make import Make


make = Make(PROJECT_DIR)
