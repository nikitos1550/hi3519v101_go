def read_till(line, pos, eot):
    result = ""
    end = len(line)
    while pos < end:
        if line[pos:].startswith(eot):
            break
        if line[pos] == "\\":
            pos += 1
        result += line[pos]
        pos += 1
    return result, pos


def read_token(line, pos, eot, strict=True, quotable=False):
    end = len(line)

    cur = line[pos]
    if quotable and cur in ("'", '"'):  # token is quoted
        pos += 1
        token, pos = read_till(line, pos, cur)

        if pos >= end:
            raise ValueError("closing quote not found")
        pos += 1

        if pos >= end:
            if strict:
                raise ValueError("EOT not found")
        elif line[pos] != eot:
            raise ValueError("EOT must be after quotes")
        pos += 1
    else:
        eot_pos = line.find(eot, pos)
        if eot_pos == -1:
            if strict:
                raise ValueError("EOT not found")
            eot_pos = end
        token = line[pos:eot_pos]
        pos = eot_pos + 1
    
    while pos < end and line[pos].isspace():  # skip spaces
        pos += 1

    return token, pos


def kv_pairs(line):
    end = len(line)
    pos = 0
    while pos < end:
        key, pos = read_token(line, pos, "=", strict=True, quotable=False)
        value, pos = read_token(line, pos, " ", strict=False, quotable=True)
        yield key.strip(), value.strip()


def kv_read(line):
    """
    KEY lasts till `=` sign and is read "as is" (no quotes or escaping are handled)
    VALUE lasts till ` ` sign but may be quoted
    """

    return {k: v for k, v in kv_pairs(line)}
