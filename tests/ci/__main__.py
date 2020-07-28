from . import PROJECT_DIR
from .utils import Git, absjoin, print_for_github_comment
from .stage import Pipeline, BrHisicamMakeAll, ApplicationMake, MakeRootFs, Deploy, CheckBuildInfo, GetBasicJpeg
from .environment import Environment
from testcore import DEVICE_LIST
import logging
import os
import time


PERSISTENT_DIR = absjoin(PROJECT_DIR, "..")
REPO_DIR = PROJECT_DIR
logging.basicConfig(level=logging.DEBUG)

e = Environment(PERSISTENT_DIR)
pipeline = Pipeline(e, stages=[
    BrHisicamMakeAll, ApplicationMake, MakeRootFs, Deploy, CheckBuildInfo, GetBasicJpeg
])

boards = DEVICE_LIST.keys()
logging.info(f"Board list: {boards}")

for board in boards:
    pipeline.run_for_board(board)
