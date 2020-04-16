from .common import DEVICE_LIST, br_hisicam, hiburn, make, PROJECT_DIR
import os
import logging
import json
import urllib.request


def br_make_and_upload(board, overlay):
    br_hisicam.make_board(board, rootfs_overlays=[overlay])

    uimage_path = br_hisicam.uimage_path(board)
    assert os.path.exists(uimage_path)

    rootfs_image_path = br_hisicam.rootfs_image_path(board)
    assert os.path.exists(rootfs_image_path)

    logging.info(f"Upload images on {board} test device and boot it...")
    info = br_hisicam.info(board)
    hiburn.boot(board, uimage=uimage_path, rootfs=rootfs_image_path, device_info=info)


def check_date(addr):
    path = "/api/system/date"
    logging.debug(f"Try {path} ...")

    resp = urllib.request.urlopen(f"http://{addr}{path}")
    date = json.loads(resp.read())

    assert "formatted" in date
    assert "secs" in date
    assert "nanosecs" in date

    logging.info(f"Handler {path} is OK")


def test_basic():
    board = "xm_53h20-s_hi3516cv100_imx122"

    info = make.info(board)
    logging.info(f"Target info:\n{info}")

    make.build_app(board)
    app_overlay = os.path.join(PROJECT_DIR, info["APP_OVERLAY"])
    logging.info(f"Application is built, overlay: {app_overlay}")

    br_make_and_upload(board, app_overlay)
    logging.info(f"Camera is running with aplication onboard")


    check_date(DEVICE_LIST[board]["ip_addr"])

