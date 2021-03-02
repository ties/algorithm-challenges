from pathlib import Path

from sherlock_and_the_valid_string import isValid

def test_sample_0():
    assert isValid("aabbcd") == "NO"


def test_sample_1():
    assert isValid("aabbccddeefghi") == "NO"


def test_sample_2():
    assert isValid("abcdefghhgfedecba") == "YES"


def test_case_3():
    assert isValid("aaaabbcc") == "NO"
