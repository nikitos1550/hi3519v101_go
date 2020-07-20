import urllib.request
from . import APPLICATION_DIR
import json
import os


PARAMS_DIR = os.path.join(APPLICATION_DIR, "api/tests")


def get_params(name):
    with open(os.path.join(PARAMS_DIR, name), "rb") as f:
        return f.read()


def get_json(url, timeout=10):
    with urllib.request.urlopen(url) as resp:
        return json.loads(resp.read().decode("utf-8"))


def post_json(url, data, timeout=10):
    return urllib.request.urlopen(
        urllib.request.Request(
            url=url,
            data=data,
            method="POST",
            headers={"Content-Type": "application/json"}
        ),
        timeout=timeout
    )


def create_channel(addr, name, params_name):
    post_json(f"http://{addr}/api/channel/{name}", get_params(params_name))

def create_encoder(addr, name, params_name):
    post_json(f"http://{addr}/api/encoder/{name}", get_params(params_name))

def link(addr, item1, item2):
    post_json(f"http://{addr}/api/link/{item1[0]}/{item1[1]}/{item2[0]}/{item2[1]}", b"")

def encoder_start(addr, name):
    get_json(f"http://{addr}/api/encoder/{name}/start")

def create_jpeg(addr, name):
    post_json(f"http://{addr}/api/jpeg/{name}", b"")


def init_basic_jpeg(addr):
    create_channel(addr, "main", "c_3840x2160.json")
    create_encoder(addr, "mjpeg_1", "e_1920x1080_mjpeg_cbr.json")
    link(addr, ("channel", "main"), ("encoder", "mjpeg_1"))
    encoder_start(addr, "mjpeg_1")
    create_jpeg(addr, "fullhd")
    link(addr, ("encoder", "mjpeg_1"), ("jpeg", "fullhd"))


def get_jpeg(addr):
    req = f"http://{addr}/serve/jpeg/fullhd.jpeg"
    print(f"REQUEST: {req}")
    with urllib.request.urlopen(req) as resp:
        data = resp.read()
    print("DATA: ", data)
    return data