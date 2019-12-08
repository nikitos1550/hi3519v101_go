import logging
import signal
import asyncio
import asyncio.subprocess
from telnetlib import Telnet


class Ser2Net:
    def __enter__(self):
        try:
            self.start()
            return self
        except:
            self.stop()
            raise

    def __exit__(self, *arg, **kwargs):
        self.stop()

    @property
    def control_port(self):
        return int(self.config["control_port"])

    @property
    def config_file(self):
        return self.config["config_file"]

    @property
    def is_running(self):
        return (self.proc is not None) and (self.proc.returncode is None)

    def __init__(self, config={}):
        self.config = config
        self.proc = None
        self.control = None

    async def start(self):
        open(self.config_file, "w").close()  # ensure that config file exists and it's clear

        args = [
            "-d",
            "-p", "localhost,{}".format(self.control_port),
            "-c", self.config_file
        ]
        logging.debug("run subprocess: ser2net {}".format(" ".join(args)))
        self.proc = await asyncio.subprocess.create_subprocess_exec("ser2net", *args, stderr=asyncio.subprocess.DEVNULL)
        logging.info("ser2net started with PID {}".format(self.proc.pid))

        await asyncio.sleep(3)
        if not self.is_running:
            raise RuntimeError("ser2net process died very fast")

        logging.debug("connect to ser2net control port...")
        self.control = Telnet(host="localhost", port=self.control_port)
        self.control.read_until(b"-> ")
        logging.info("connection with ser2net control port established")

    async def stop(self):
        if not self.is_running:
            return

        logging.debug("ser2net process is stopping...")
        self.proc.terminate()
        try:
            await asyncio.wait_for(self.proc.wait(), timeout=5)
            logging.info("ser2net process successfully terminated")
        except asyncio.TimeoutError:
            self.proc.kill()
            await self.proc.wait()
            logging.warn("ser2net process killed")

    def reload_config(self):
        self.proc.send_signal(signal.SIGHUP)

    def control_cmd(self, command):
        req = command.encode("ascii") + b"\n\r"
        logging.debug("request to ser2net control: {}".format(req))
        self.control.write(req)
        res = self.control.read_until(b"-> ")
        logging.debug("response from ser2net control: {}".format(req))
        return res[len(data):-5].decode("ascii")

    def disconnect(self, port):
        resp = self.control_cmd("disconnect {}".format(port))
        if resp != "":
            raise RuntimeError("disconnect failed: {}".format(resp))

    def setportenable(self, port, state="telnet"):
        if state not in ("raw", "rawIp", "telnet", "off"):
            raise ValueError("'state' argument has inappropriate value")
        resp = self.control_cmd("setportenable {} {}".format(port, state))
        if resp != "":
            raise RuntimeError("setportenable failed: {}".format(resp))

    def showport(self, port):
        resp = self.control_cmd("showshortport {}".format(port))
        vals = filter(None, resp.split("\n")[1].strip().split(" "))

        fields = ("port", "type", "timeout", "remote_addr", "device")
        result = {}
        for f in fields:
            result[f] = vals.__next__()
        return result
