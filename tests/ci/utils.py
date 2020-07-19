from . import PROJECT_DIR
import subprocess
import json
import time
import os
import shutil


def read_file(path):
    if not os.path.exists(path):
        return None
    with open(path, "r") as f:
        return f.read()


def absjoin(*args):
    return os.path.abspath(os.path.join(*args))


def rmdir(path):
    if os.path.exists(path):
        shutil.rmtree(path)


def copydir(src, dst):
    rmdir(dst)
    shutil.copytree(src, dst)


def request_json(url, timeout=10):
    import urllib.request

    deadline = time.monotonic() + timeout
    while True:
        try:
            with urllib.request.urlopen(url) as resp:
                return json.loads(resp.read().decode("utf-8"))
        except:
            if time.monotonic() > deadline:
                raise


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
