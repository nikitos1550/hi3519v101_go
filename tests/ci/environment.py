from testcore.brhisicam import BrHisiCam
from testcore.make import Make
from .utils import Git, read_file, absjoin
from .stage import Stage
from . import CI_DIR, APPLICATION_DIR, BR_HISICAM_DIR
import os
import sys
import time
import subprocess
import shutil
import logging
import logging.handlers


#   PERSISTENT_DIR
#   +-- br-hisicam_gitref   - contains gitref of last used br-hisicam
#   +-- br-hisicam_output   - artifacts of last br-hisicam
#   |   +-- <BOARD-ID>      - contains output of br-hisicam for particular board
#   |   +-- ...
#   +-- runs
#       +-- <DATE>_<REPO_GITREF>        - contains artifacts of particular CI run
#       |   +-- go_workdir              - GOPATH
#       |   +-- <BOARD-ID>              - artifacts of particular board
#       |   |   +-- params.mk           - parameters of the board
#       |   |   +-- app_outdir          - application's artifacts
#       |   |   |   +-- overlay         - application's overlay directory (distrib)
#       |   |   |   +-- ...
#       |   |   +-- logs
#       |   |       +-- common.log      - common log for this CI run
#       |   |       +-- <STAGE-ID>.log  - log of particular CI stage
#       |   |       +-- ...
#       |   +-- ...
#       +-- ...


BR_HISICAM_OUTDIR_ROOT = "br-hisicam_output"
BR_HISICAM_GITREF_FILE = "br-hisicam_gitref"
APP_OUTDIR = os.path.join(CI_DIR, "app_outdir")
GO_PATH = os.path.join(CI_DIR, "go_workdir")


def prepare_br_hisicam_outdir(persistent_dir):
    outdir = absjoin(persistent_dir, BR_HISICAM_OUTDIR_ROOT)
    gitref_file = absjoin(persistent_dir, BR_HISICAM_GITREF_FILE)
    current_gitref = Git.get_submodule_ref("br-hisicam")
    existing_gitref = read_file(gitref_file)
    
    logging.info(f"BR gitref current={current_gitref}, existing={existing_gitref}")

    if current_gitref != existing_gitref:
        if os.path.exists(outdir):
            logging.info(f"Remove existing BR output root directory '{outdir}'")
            shutil.rmtree(outdir)

    os.makedirs(outdir, exist_ok=True)
    with open(gitref_file, "w") as f:
        f.write(current_gitref)

    logging.info(f"BR output directory is ready: {outdir}")
    return outdir


def prepare_rundir(persistent_dir):
    gitref = Git.get_ref()
    date = time.strftime("%Y%m%dT%H%M%S", time.localtime())

    rundir = absjoin(persistent_dir, f"runs/{date}_{gitref}")
    if os.path.exists(rundir):
        shutil.rmtree(rundir)

    os.makedirs(rundir)
    return rundir


class Environment:
    def __init__(self, persistent_dir):
        self.br_hisicam_outdir_root = prepare_br_hisicam_outdir(persistent_dir)
        self.rundir_root = prepare_rundir(persistent_dir)
        self.go_workdir = absjoin(self.rundir_root, "go_workdir")
        self.app_outdir_root = absjoin(self.rundir_root, APP_OUTDIR)

    def get_stage_output_files(self, stage_name, board):
        logsdir = absjoin(self.rundir_root, board, "logs")
        os.makedirs(logsdir, exist_ok=True)
        base_fname = absjoin(logsdir, stage_name)
        return f"{base_fname}.out", f"{base_fname}.err"

    def board_params_file(self, board):
        return absjoin(self.rundir_root, board, "params.mk")
    
    def app_output_dir(self, board):
        return absjoin(self.rundir_root, board, "app_outdir")
    
    def app_overlay_dir(self, board):
        return absjoin(self.rundir_root, board, "app_outdir/overlay")


class ApplicationMake(Stage):
    def run(self, board):
        params_file = self.env.br_hisicam_params_file(board)
        outdir = self.env.app_outdir(board)

        make = Make(
            root_dir=APPLICATION_DIR,
            args=[f"PARAMS_FILE={params_file}", f"OUTDIR={outdir}", f"GOPATH={self.env.go_workdir}"],
            stdout=sys.stderr.fileno(),
            stderr=None
        )

        logging.info(f"Build application for {board} board...")
        make.check_call(["build-cam"])


class CreateRootFS(Stage):
    def run(self, board):
        outdir = self.env.app_outdir(board)

        br_hisicam = BrHisiCam(board=board, output_dir=outdir)

        logging.info(f"Get info of {board} board...")
        info = br_hisicam.make_board_info()

        overlay_dir = absjoin(APPLICATION_DIR, "distrib", info["FAMILY"])
        logging.info(f"Application overlay dir: {overlay_dir}")

        logging.info(f"Create overlayed root FS for {board} board...")
        br_hisicam.make_overlayed_rootfs(overlays=[overlay_dir])
