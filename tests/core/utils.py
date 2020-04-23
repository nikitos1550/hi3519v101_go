
def parse_kv_lines(lines):
    result = {}
    for line in lines:
        if not line:
            continue
        try:
            key, value = line.split("=")
            result[key.strip()] = value.strip()
        except ValueError as err:
            logging.warning(f"Could not unpack line '{line}' - {err}")
    return result
