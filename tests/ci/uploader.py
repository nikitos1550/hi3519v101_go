import hashlib
import hmac
import time
import os
import urllib.request


KEY_ID = os.getenv("S3_KEY_ID")
KEY = os.getenv("S3_KEY")


def sha256hash(data):
    if isinstance(data, str):
        data = data.encode("utf-8")
    m = hashlib.sha256()
    m.update(data)
    return m.digest()


def sign(key, data):
    if isinstance(key, str):
        key = key.encode("utf-8")
    if isinstance(data, str):
        data = data.encode("utf-8")
    return hmac.digest(key, data, "sha256")


def get_key(secret, date):
    return sign(sign(sign(sign("AWS4" + secret, date), "ru-central1"), "s3"), "aws4_request")


_signed_headers = ("host", "x-amz-content-sha256", "x-amz-date")


def sign_request(method, uri, host):
    now = time.gmtime()
    timestamp = time.strftime("%Y%m%dT%H%M%SZ", now)
    date = time.strftime("%Y%m%d", now)
    scope = f"{date}/ru-central1/s3/aws4_request"

    headers = {
        "host": host,
        "x-amz-content-sha256": "UNSIGNED-PAYLOAD",
        "x-amz-date": timestamp
    }

    # make canonical request
    canon = (
        method.upper() + "\n" +
        uri + "\n" +
        "" + "\n"  # no query params
    )
    for h in _signed_headers:
        canon += h + ":" + headers[h] + "\n"
    canon += "\n"
    canon += ";".join(_signed_headers) + "\n"
    canon += "UNSIGNED-PAYLOAD"

    canon_hash = sha256hash(canon.encode("utf-8")).hex()

    data = f"AWS4-HMAC-SHA256\n{timestamp}\n{scope}\n{canon_hash}"

    auth_value = (
        "AWS4-HMAC-SHA256 "
        f"Credential={KEY_ID}/{date}/ru-central1/s3/aws4_request,"
        f"SignedHeaders={';'.join(_signed_headers)},"
        f"Signature={sign(get_key(KEY, date), data).hex()}"
    )

    headers["Authorization"] = auth_value
    return headers


def make_s3_request(method, uri, host, headers={}, **kwargs):
    headers.update(
        sign_request(method=method, uri=uri, host=host)
    )
    return urllib.request.Request(
        url=f"https://{host}{uri}",
        headers=headers,
        method=method,
        **kwargs
    )

# -------------------------------------------------------------------------------------------------


def upload_jpeg_image(data, key):
    HOST = "storage.yandexcloud.net"
    req = make_s3_request(
        method="PUT",
        uri=f"/larvaci-hisicam/{key}",
        host=HOST,
        headers={
            "Content-Type": "image/jpeg"
        },
        data=data
    )
    resp = urllib.request.urlopen(req)
    return f"https://{HOST}/larvaci-hisicam/{key}"



if __name__ == "__main__":
    upload_jpeg_image("./basic.jpeg", "pics/basic.jpeg")