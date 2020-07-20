from . import CI_DIR, BR_HISICAM_DIR, APPLICATION_DIR
from .utils import copydir, absjoin, request_json
from testcore.brhisicam import BrHisiCam
from testcore.make import Make
from testcore import hiburn
import logging
import os
import sys
import subprocess
import time
import json


class Stage:
    def __init__(self, env, board, pipeline):
        self.name = self.__class__.__name__
        self.env = env
        self.board = board
        self.pipeline = pipeline

        self.logger = logging.getLogger(f"{self.board}/{self.name}")
        
        self.br_hisicam_outdir = os.path.join(self.env.br_hisicam_outdir_root, self.board)
        self.app_outdir = self.env.app_output_dir(self.board)
        self.app_overlay_dir = self.env.app_overlay_dir(self.board)
        self.go_workdir = self.env.go_workdir
        self.params_file = self.env.board_params_file(board)
        self.stdout, self.stderr = self.env.get_stage_output_files(self.name, self.board)

        self.br_hisicam = BrHisiCam(
            board=self.board,
            output_dir=self.br_hisicam_outdir,
            stdout=self.stdout,
            stderr=self.stdout  # to keep all output in order
        )

    @property
    def uimage_path(self):
        return absjoin(self.br_hisicam_outdir, "images/uImage")
    
    @property
    def overlayed_rootfs_path(self):
        return absjoin(self.br_hisicam_outdir, "images/rootfs-overlayed.squashfs")

    def run(self, *args, **kwargs):
        raise NotImplementedError("Method 'run' must be defined in subclasses")

    def debug(self, *args, **kwargs):
        self.logger.debug(*args, **kwargs)

    def info(self, *args, **kwargs):
        self.logger.info(*args, **kwargs)

    def warning(self, *args, **kwargs):
        self.logger.warning(*args, **kwargs)

    def error(self, *args, **kwargs):
        self.logger.error(*args, **kwargs)

    def exception(self, *args, **kwargs):
        self.logger.exception(*args, **kwargs)
    
    def state(self, msg):
        self.pipeline.set_stage_state(self, msg)


# -------------------------------------------------------------------------------------------------
class Pipeline:
    def __init__(self, env, stages):
        self.env = env
        self.stages = stages
        self.boards = []

    def make_report(self):
        report = ""
        report += " Board |" + "|".join(f" {s.__name__} " for s in self.stages) + "\n"
        report += "-------|" + "|".join("-" * (len(s.__name__) + 2) for s in self.stages) + "\n"
        for b in self.boards:
            report += f" {b[0]} |" + "|".join(b[1].get(s.__name__, "") for s in self.stages) + "\n"
        return report

    def _get_states(self, board):
        for b in self.boards:
            if board == b[0]:
                return b[1]

    def print_report(self):
        sys.stdout.buffer.write(self.make_report().encode("utf-8") + b"\x04")
        sys.stdout.flush()
    
    def state_set(self, stage, msg):
        states = self._get_states(stage.board)
        states[stage.name] = msg
        self.print_report()
    
    def state_add(self, stage, msg):
        states = self._get_states(stage.board)
        states[stage.name] = states[stage.name] + msg
        self.print_report()
    
    def run_for_board(self, board):
        logging.info(f"Start pipeline for board '{board}'...")
        self.boards.append((board, {}))
        for cls in self.stages:
            stage = cls(self.env, board, self)
            try:
                logging.info(f"Start stage '{stage.name}' for board '{board}'...")
                self.state_set(stage, "started...")
                stage.run()
                self.state_add(stage, " (OK)")
                logging.info(f"Stage '{stage.name}' successfully finished for board '{board}'")
            except AssertionError as err:
                logging.exception(f"Stage '{stage.name}' failed with assertion error for board '{board}'")
                self.state_add(f" (failed: {err})")
                return
            except:
                logging.exception(f"Stage '{stage.name}' failed with exception for board '{board}'")
                self.state_add(stage, " (failed with exception)")
                return


# -------------------------------------------------------------------------------------------------
class BrHisicamMakeAll(Stage):
    def run(self):
        CREATE_PARAMS_SCRIPT = os.path.join(CI_DIR, "create-params-file.sh")

        self.info("Make br-hisicam...")
        self.br_hisicam.make_all("opus", "host-go", "toolchain-params")
        
        self.info("Create board params file...")
        subprocess.check_call(
            f"{CREATE_PARAMS_SCRIPT} {self.board} {self.br_hisicam_outdir} {BR_HISICAM_DIR} > {self.params_file}",
            shell=True
        )


# -------------------------------------------------------------------------------------------------
class ApplicationMake(Stage):
    def run(self):
        make = Make(
            root_dir=APPLICATION_DIR,
            args=[f"PARAMS_FILE={self.params_file}", f"OUTDIR={self.app_outdir}", f"GOPATH={self.go_workdir}"],
            stdout=self.stdout,
            stderr=self.stdout,
            verbose=True
        )

        logging.info(f"Clean...")
        make.check_call(["clean"])

        logging.info(f"Build...")
        make.check_call(["build-cam"])

        logging.info(f"Get board info...")
        info = self.br_hisicam.make_board_info()

        logging.info(f"Copy overlay directory... ")
        copydir(
            src=os.path.join(APPLICATION_DIR, "distrib", info["FAMILY"]),
            dst=self.app_overlay_dir
        )


# -------------------------------------------------------------------------------------------------
class MakeRootFs(Stage):
    def run(self):
        logging.info(f"Create overlayed root FS...")
        self.br_hisicam.make_overlayed_rootfs(overlays=[self.app_overlay_dir])


# -------------------------------------------------------------------------------------------------
class Deploy(Stage):
    def run(self):
        logging.info(f"Get board info...")
        info = self.br_hisicam.make_board_info()

        logging.info(f"Deploy on device...")
        with open(self.stdout, "ab") as fout:
            hiburn.boot(
                self.board,
                uimage=self.uimage_path,
                rootfs=self.overlayed_rootfs_path,
                device_info=info,
                stdout=fout, stderr=fout,
                timeout=180
            )


# -------------------------------------------------------------------------------------------------
class CheckBuildInfo(Stage):
    def run(self):
        from testcore import DEVICE_LIST

        addr = DEVICE_LIST[self.board]["ip_addr"]
        url = f"http://{addr}/api/buildinfo"
        
        resp = request_json(url, timeout=20)
        logging.info(f"Got build info: {json.dumps(resp, indent=4)}")

        buildcommit = resp["buildcommit"].strip()
        assert buildcommit == self.env.gitref, "Invalid buildinfo"