from .utils import parse_kv_lines
import subprocess


class Make:
    def __init__(self, root_dir):
        self.root_dir = root_dir

    def info(self, board):
        out = subprocess.check_output([
            "make", "NO_USER_MAKEFILE=Y", f"BOARD={board}", "APP_TARGET=cam", "info"
        ], cwd=self.root_dir)
        return parse_kv_lines(out.decode("utf-8").split("\n"))

    def build_app(self, board):
        subprocess.check_call([
            "make", "NO_USER_MAKEFILE=Y", f"BOARD={board}", "APP_TARGET=cam", "build-app"
        ], cwd=self.root_dir)
