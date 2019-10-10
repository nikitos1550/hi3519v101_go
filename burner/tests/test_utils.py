from .. import utils
from ..kv_read import kv_read
import tempfile


def test_from_hsize():
    assert utils.from_hsize("14") == utils.from_hsize("14B") == utils.from_hsize("14b") == 14
    assert utils.from_hsize("14k") == utils.from_hsize("14K") == 14 * 1024
    assert utils.from_hsize("14m") == utils.from_hsize("14M") == 14 * 1024  *1024
    assert utils.from_hsize("14g") == utils.from_hsize("14G") == 14 * 1024 * 1024  *1024


def test_to_hsize():
    assert utils.to_hsize(0) == "0"
    assert utils.to_hsize(1) == "1"
    assert utils.to_hsize(1023) == "1023"
    assert utils.to_hsize(1024) == "1K"
    assert utils.to_hsize(1025) == "1025"
    assert utils.to_hsize(1024 * 19) == "19K"
    assert utils.to_hsize(1024 * 1024 * 4) == "4M"
    assert utils.to_hsize(1024 * 1024 * 1024 * 4) == "4G"


def test_aligned_address():
    assert utils.aligned_address(32, 0) == 0
    assert utils.aligned_address(32, 1) == 32
    assert utils.aligned_address(32, 31) == 32
    assert utils.aligned_address(32, 32) == 32
    assert utils.aligned_address(32, 33) == 64


def test_read_parameters_from_file():
    with tempfile.NamedTemporaryFile(mode="wt") as f:
        f.write(
            "\nPARAM_1 = 12345"
            "\n   PARAM_2=12345    "
            "\n PARAM_2  =  54321  "  # overrides previous param
            "\n  # PARAM_3 = it's a comment actually \\"
            "\n  \n\n \t \t \n"  # some empty lines
            "\nPARAM WITH SPACES = the last value"
            "\n# comment again"
        )
        f.flush()

        assert utils.read_parameters_from_file(f.name) == {
            "PARAM_1": "12345",
            "PARAM_2": "54321",
            "PARAM WITH SPACES": "the last value"
        }


def test_kv_read():
    assert kv_read("A=1 B = 2  C  =  3   ") == {"A": "1", "B": "2", "C": "3"}
    assert kv_read("TheKey=TheLongValue   The Key= 'The Long Value'   \"The'Key\"= '\"The\\'Long\\'Value\"'") == {
        "TheKey": "TheLongValue",
        "The Key": "The Long Value",
        "\"The'Key\"": "\"The'Long'Value\""
    }
