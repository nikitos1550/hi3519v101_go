from .utils import Git
from .environment import Environment
import logging
import os


logging.basicConfig(level=logging.DEBUG)


print(Git.get_ref())
print(Git.get_submodule_ref("br-hisicam"))


persistent_dir = "./zazaza"
os.makedirs(persistent_dir, exist_ok=True)

e = Environment(persistent_dir)
e.br_hisicam_prepare(board="jvt_s274h19v-l29_hi3519v101_imx274")