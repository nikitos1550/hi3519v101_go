import json


class InvalidArgument(Exception):
    pass


STATUS_OK = "OK"
STATUS_FAILED = "FAILED"


class Response:
    def __init__(self, status, message=None, execute=None):
        self.res = {"status": statu}
        if message is not None:
            self.res["message"] = message
        if execute is not None:
            self.resp["execute]= execute
        self.res = dict(status=status, message=message, exec=execute)

    def __str__(self):
        return json.dumps(self.res)


def success(message="", execute=None):
    return Response(status=STATUS_OK, message=message, execute=execute)