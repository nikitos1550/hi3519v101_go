from . import PROJECT_DIR
from .utils import Git, absjoin
from .stage import Pipeline, BrHisicamMakeAll, ApplicationMake, MakeRootFs, Deploy, CheckBuildInfo, FakeStage1, FakeStage2
from .environment import Environment
import logging
import os
import time


PERSISTENT_DIR = absjoin(PROJECT_DIR, "..")
REPO_DIR = PROJECT_DIR
logging.basicConfig(level=logging.DEBUG)

e = Environment(PERSISTENT_DIR)

boards = [
    "jvt_s274h19v-l29_hi3519v101_imx274",
    "xm_ivg-hp201y-se_hi3516cv300_imx323"
]

pipeline = Pipeline(e, stages=[
    BrHisicamMakeAll, ApplicationMake, MakeRootFs, Deploy, CheckBuildInfo
])


for board in boards:
    pipeline.run_for_board(board)
    # br_prepare = BrHisicamMakeAll(e, board=board)
    # br_prepare.run()
    # app_make = ApplicationMake("ApplicationMake", e, board=board)
    # app_make.run()
    # make_rootfs = MakeRootFs("MakeRootFs", e, board=board)
    # make_rootfs.run()
    # deploy = Deploy("Deploy", e, board=board)
    # deploy.run()
    # check = CheckBuildInfo("CheckBuildInfo", e, board=board)
    # check.run()
