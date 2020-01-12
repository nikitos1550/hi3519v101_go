#!/bin/env python3
import socket
import json
import os
from . import common


class FailedRequest(Exception):
    def __init__(self, response):
        self.response = response


class Connection:
    def __init__(self, host, port):
        self.sock = socket.create_connection(address=(host, port))
        self._buf = bytearray()
        self._readuntil(b"#")

    def close(self):
        self.sock.close()

    def _readuntil(self, end):
        res = bytearray()
        while True:
            while len(self._buf):
                res.append(self._buf.pop(0))
                if res[-len(end):] == end:
                    return res
            self._buf += self.sock.recv(4096)

    def _write(self, data):
        self.sock.sendall(data)

    def _make_request(self, request):
        self._write(request.encode("ascii") + b"\n")
        response = self._readuntil(b"#")
        return response[:-2].decode("ascii")

    def request(self, req):
        resp = json.loads(self._make_request(req))
        if resp["status"] != common.STATUS_OK:
            raise FailedRequest(resp)
        return resp


class Client:
    def __init__(self, host, port):
        self.conn = Connection(host, port)
        self._set_user()

    def _set_user(self):
        self.conn.request("set_user {}".format(os.getlogin()))
    
    def list_devices(self):
        """ Get list of available devices as pairs (<name>,<model>)
        """
        devices = []
        resp = self.conn.request("list_devices")
        for l in resp["message"].split("\n"):
            devices.append(l.split(" "))
        return devices

    def forward_serial(self, devname):
        """ Request to forward serial port via TCP (like Telnet)
        Returns TCP endpoint as (host, port)
        """
        resp = self.conn.request(f"forward_serial {devname}")
        _, host, port = resp["exec"].split(" ")
        return (host, port)
    
    def release_device(self, devname):
        """ Release previously acquired device
        """
        self.conn.request(f"release_device {devname}")

    def acquire_device(self, devname):
        """ Acquire device for exclusive usage (or refresh its' tmieout)
        """
        self.conn.request(f"acquire_device {devname}")





if __name__ == "__main__":
    cl = Client("localhost", 43500)
    print(cl.forward_serial("/dev/ttyCAM1"))
    print(cl.release_device("/dev/ttyCAM2"))
