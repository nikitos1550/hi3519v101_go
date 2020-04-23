from .core import DEVICE_LIST, br_hisicam, hiburn, make, PROJECT_DIR
from .core import checks
import os
import logging
import json
import urllib.request


def br_make_and_upload(board, overlay):
    br_hisicam.make_board(board, rootfs_overlays=[overlay])

    uimage_path = br_hisicam.uimage_path(board)
    rootfs_image_path = br_hisicam.rootfs_image_path(board)

    logging.info(f"Upload images on {board} test device and boot it...")
    info = br_hisicam.info(board)
    hiburn.boot(board, uimage=uimage_path, rootfs=rootfs_image_path, device_info=info)


def test_basic():
    board = "jvt_s274h19v-l29_hi3519v101_imx274"

    info = make.info(board)
    logging.info(f"Target info:\n{info}")

    make.build_app(board)
    app_overlay = os.path.join(PROJECT_DIR, info["APP_OVERLAY"])
    logging.info(f"Application is built, overlay: {app_overlay}")

    br_make_and_upload(board, app_overlay)
    logging.info(f"Camera is running with aplication onboard")

    failed = checks.run_checks(DEVICE_LIST[board]["ip_addr"])
    assert bool(failed) is False
