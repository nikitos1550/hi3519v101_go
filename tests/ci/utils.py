from . import PROJECT_DIR
import subprocess
import os


def read_file(path):
    if not os.path.exists(path):
        return None
    with open(path, "r") as f:
        return f.read()


class Git:
    GIT_BINARY = "git"
    
    @classmethod
    def get_ref(cls, repodir=None):
        if repodir is None:
            repodir = PROJECT_DIR
        out = subprocess.check_output([cls.GIT_BINARY, "rev-parse", "HEAD"], cwd=repodir)
        return out.decode("utf-8").strip()
    
    @classmethod
    def get_submodule_ref(cls, submodule):
        return cls.get_ref(repodir=os.path.join(PROJECT_DIR, submodule))
