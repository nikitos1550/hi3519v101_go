import os
import sys


CI_DIR                      = os.path.dirname(__file__)
PROJECT_DIR                 = os.path.abspath(os.path.join(CI_DIR, "../.."))
APPLICATION_DIR             = os.path.abspath(os.path.join(PROJECT_DIR, "application"))
BR_HISICAM_DIR              = os.path.abspath(os.path.join(PROJECT_DIR, "br-hisicam"))
BR_HISICAM_TESTCORE_DIR     = os.path.abspath(os.path.join(BR_HISICAM_DIR, "testcore"))


if not os.path.isdir(BR_HISICAM_TESTCORE_DIR):
    raise SystemExit(
        f"BR_HISICAM_TESTCORE_DIR directory ({BR_HISICAM_TESTCORE_DIR}) is absent"
    )

sys.path.insert(0, BR_HISICAM_DIR)
