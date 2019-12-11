import subprocess
import os
import time
import pytest


CAMSTORE_SH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../camstore.sh"))
CAMSTORE_PORT = 56342
CLIENT_MAIN = os.path.abspath(os.path.join(os.path.dirname(__file__), "../client/main.py"))



@pytest.fixture()
def proc():
    with subprocess.Popen([CAMSTORE_SH, "start_testing"], stdout=subprocess.PIPE, stderr=subprocess.PIPE) as proc:
        time.sleep(2)
        yield proc


def test_run(proc):
    completed = subprocess.run([CLIENT_MAIN, "--port", str(CAMSTORE_PORT), "help"], check=True)
    print(completed.stdout)

    assert proc.returncode is None  # process is still alive
    proc.terminate()
    proc.wait(timeout=5)
    assert proc.returncode is not None