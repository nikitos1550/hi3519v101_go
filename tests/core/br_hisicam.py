BR_HISICAM_TESTENV_DIR = os.path.abspath(os.path.join(PROJECT_DIR, "br-hisicam/tests"))


if not os.path.isdir(BR_HISICAM_TESTENV_DIR):
    raise SystemExit(
        "br-hisicam/tests directory is absent. Make sure 'br-hisicam' submodule is initialized"
    )

sys.path.insert(0, BR_HISICAM_TESTENV_DIR)


from testenv import DEVICE_LIST, br_hisicam, hiburn

# -------------------------------------------------------------------------------------------------


def br_make_and_upload(board, overlay):
    br_hisicam.make_board(board, rootfs_overlays=[overlay])

    uimage_path = br_hisicam.uimage_path(board)
    assert os.path.exists(uimage_path)

    rootfs_image_path = br_hisicam.rootfs_image_path(board)
    assert os.path.exists(rootfs_image_path)

    logging.info(f"Upload images on {board} test device and boot it...")
    info = br_hisicam.info(board)
    hiburn.boot(board, uimage=uimage_path, rootfs=rootfs_image_path, device_info=info)
