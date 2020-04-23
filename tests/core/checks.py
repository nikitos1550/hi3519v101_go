#!/usr/bin/env python3
import logging
import json
import urllib.request
import functools


ALL_CHECKS = {}


def register_check(func):
    name = func.__name__
    if name in ALL_CHECKS:
        raise RuntimeError(f"Check with name '{name}' already exists")

    @functools.wraps(func)
    def wrapped(*args, **kwargs):
        logging.info(f"Run check '{name}'...")
        try:
            ret = func(*args, **kwargs)
            logging.info(f"Check '{name}' successfully passed")
            return ret
        except:
            logging.error(f"Check '{name}' failed")
            raise

    ALL_CHECKS[name] = wrapped
    return wrapped


def get_json(url):
    resp = urllib.request.urlopen(url)
    jresp = json.loads(resp.read())
    logging.debug(f"Made HTTP request '{url}'; response: {jresp}")
    return jresp


def list_checks():
    for name, func in ALL_CHECKS.items():
        doc = func.__doc__
        if doc is None:
            yield f" - {name}"
        if doc is not None:
            desc = doc.strip().split("\n")[0].strip()
            yield f" - {name} - {desc}"


def run_checks(endpoint, check_list=None):
    if check_list is None:
        check_list = (name for name in ALL_CHECKS.keys())

    failed = {}
    for name in check_list:
        check = ALL_CHECKS[name]
        try:
            check(endpoint)
        except Exception as err:
            failed[name] = err

    return failed


# -------------------------------------------------------------------------------------------------
@register_check
def check_date(endpoint):
    """ Check "/api/system/date" handle; it should return current date
    """

    date = get_json(f"http://{endpoint}/api/system/d1ate")

    assert "formatted" in date
    assert "secs" in date
    assert "nanosecs" in date


@register_check
def check_umaps(endpoint):
    umaps = get_json(f"http://{endpoint}/api/debug/umap.json")

    logging.info(umaps)

    for umap in umaps:
        umapjson = get_json(f"http://{endpoint}/api/debug/umap/{umap}.json")

    logging.info(f"Handler umaps is OK")


@register_check
def check_temperature(endpoint):
    """ Check '/api/temperature' handle; it should return current temperature
    """
    temperature = get_json(f"http://{endpoint}/api/temperature")
    logging.info(f"Handler {temperature} is OK")


@register_check
def check_mpp(addr):
    path = "/api/mpp"


# -------------------------------------------------------------------------------------------------
def main():
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument("--verbose", "-v", action="store_true", help="Print debug output")
    parser.add_argument("--list", "-l", action="store_true", help="List available checks")
    parser.add_argument("--endpoint", "-e", type=str, help="Service's HTTP endpoint (host[:port])")
    parser.add_argument("checks", type=str, nargs="*", help="Checks to run")

    args = parser.parse_args()

    logging.basicConfig(level=(logging.DEBUG if args.verbose else logging.INFO))

    if args.list:
        print("\n".join(list_checks()))
        exit(0)
    if args.endpoint is None:
        print("Endpoint MUST be set")
        exit(1)

    failed = run_checks(endpoint=args.endpoint, check_list=(args.checks or None))
    if not failed:
        print("ALL PASSED")
        exit(0)

    print("FAILED CHECKS:")
    for name, err in failed.items():
        print(f"- {name}: {err}")
    exit(1)


if __name__ == "__main__":
    main()