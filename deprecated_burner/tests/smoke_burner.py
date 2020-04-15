import subprocess
import os
import pytest


devices = [
    {"num": 1, "port": "/dev/ttyCAM1"},
    {"num": 2, "port": "/dev/ttyCAM2"},
    {"num": 5, "port": "/dev/ttyCAM5", "uboot_params": "PROMPT='Zview #'"},
    {"num": 6, "port": "/dev/ttyCAM6"},
    {"num": 7, "port": "/dev/ttyCAM7", "uboot_params": "PROMPT='hisilicon #' GREETING=U-Boot"},
    {"num": 8, "port": "/dev/ttyCAM8", "uboot_params": "PROMPT='xmtech #' GREETING=U-Boot"},
    {"num": 9, "port": "/dev/ttyCAM9", "uboot_params": "PROMPT='Zview #'"},
    {"num": 10, "port": "/dev/ttyCAM10", "uboot_params": "PROMPT='Zview #'"},
    {"num": 12, "port": "/dev/ttyCAM12", "uboot_params": "PROMPT='hisilicon #' GREETING=U-Boot"},
    {"num": 13, "port": "/dev/ttyCAM13", "uboot_params": "PROMPT='xmtech #' GREETING='System startup'"},
    {"num": 14, "port": "/dev/ttyCAM14", "uboot_params": "PROMPT='hisilicon #' GREETING=U-Boot AUTOBOOT_STOP_KEY=1"},

    
]


BURNER_WORKDIR = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))

def launch(params):
    args = [
        "./burner2.py",
        "--reset-power", "./power.py reset {}".format(params["num"]),
        "--port", params["port"],
        "--uboot-params", params.get("uboot_params", ""),
        "printenv"
    ]
    subprocess.check_call(args, cwd=BURNER_WORKDIR)


@pytest.mark.parametrize("params", devices)
def test_print_env(params):
    launch(params)
