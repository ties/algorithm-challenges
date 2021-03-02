from special_substring import substrCount, drop_middle
from pathlib import Path

def test_drop_middle():
    assert drop_middle("") == ""
    assert drop_middle("a") == "a"
    assert drop_middle("ab") == "ab"
    assert drop_middle("abc") == "ac"
    assert drop_middle("abcde") == "abde"

def test_sample_in_text():
    assert substrCount(8, "mnopopoo") == 12

def test_sample_0():
    assert substrCount(5, "asasd") == 7

def test_sample_1():
    assert substrCount(7, "abcbaba") == 10

def test_sample_2():
    assert substrCount(4, "aaaa") == 10


def test_case_2():
    root = Path(__file__).parent
    input02 = (root /"special_substring_input02.txt").open().readlines()

    assert substrCount(int(input02[0].strip()), input02[1].strip()) == 1272919
