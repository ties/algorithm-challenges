from pathlib import Path

from alternating_characters import alternatingCharacters


def test_sample_0():
    assert alternatingCharacters("AAAA") == 3

    assert alternatingCharacters("BBBBB") == 4
    assert alternatingCharacters("ABABABAB") == 0
    assert alternatingCharacters("BABABA") == 0
    assert alternatingCharacters("AAABBB") == 4


def test_sample_1():
    assert alternatingCharacters("AAABBBAABB") == 6
    assert alternatingCharacters("AABBAABB") == 4
    assert alternatingCharacters("BABABAA") == 1


def test_sample_2():
    assert alternatingCharacters("ABBABBAA") == 3

def test_case_4():
    root = Path(__file__).parent
    outputs = (root /"alternating_characters_test_case_4_output.txt").open().readlines()

    with (root/"alternating_characters_test_case_4.txt").open() as f:
        for idx, line in enumerate(f.readlines()):
            if idx == 0:
                continue

            assert alternatingCharacters(line.strip()) == int(outputs[idx-1].strip())

