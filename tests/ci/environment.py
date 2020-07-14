from testcore.brhisicam import BrHisiCam
from testcore.make import Make
from .utils import Git, read_file
from . import CI_DIR, APPLICATION_DIR, BR_HISICAM_DIR
import os
import sys
import subprocess
import shutil
import logging



_BR_HISICAM_OUTDIR_ROOT = "br-hisicam_outdir"
_BR_HISICAM_GITREF_FILE = ".br-hisicam_git-ref"
_CREATE_PARAMS_SCRIPT = os.path.join(CI_DIR, "create-params-file.sh")


class Environment:
    def __init__(self, persistent_dir):
        self.br_hisicam_outdir_root = os.path.join(persistent_dir, _BR_HISICAM_OUTDIR_ROOT)
        self.br_hisicam_gitref_file = os.path.join(persistent_dir, _BR_HISICAM_GITREF_FILE)
        self.br_hisicam_gitref = Git.get_submodule_ref("br-hisicam")

        if self.br_hisicam_gitref != read_file(self.br_hisicam_gitref_file):
            if os.path.exists(self.br_hisicam_outdir_root):
                logging.info(f"Remove existing BR output root directory '{self.br_hisicam_outdir_root}'")
                shutil.rmtree(self.br_hisicam_outdir_root)

        os.makedirs(self.br_hisicam_outdir_root, exist_ok=True)
        with open(self.br_hisicam_gitref_file, "w") as f:
            f.write(self.br_hisicam_gitref)

    def br_hisicam_prepare(self, board):
        board_outdir = os.path.join(self.br_hisicam_outdir_root, board)
        board_params = os.path.join(self.br_hisicam_outdir_root, f"{board}-params.mk")

        br_hisicam = BrHisiCam(
            board=board,
            output_dir=board_outdir
        )
        br_hisicam.make_all("opus", "host-go")

        subprocess.check_call(
            f"{_CREATE_PARAMS_SCRIPT} {board} {board_outdir} {BR_HISICAM_DIR} > {board_params}",
            shell=True
        )

        return board_params

    def application_make(self, board):
        make = Make(root_dir=APPLICATION_DIR, args=[],
            stdout=sys.stderr.fileno(),
            stderr=None):

        params_file = self.br_hisicam_prepare(board)

