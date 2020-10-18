from typing import Counter, List
from collections import Counter


def shortest_prefix(xs: List[str]):
    counter: Counter = Counter()
    for x in xs:
        for i, c in enumerate(x):
            counter["{}{}".format(i, c)] += 1

    out: List[str] = []
    for x in xs:
        prefix = ""
        for i, c in enumerate(x):
            prefix += c
            if (counter["{}{}".format(i, c)]) == 1:
                out.append(prefix)
                break
    return out


def test_answer():
    input = ["dog", "cat", "apple", "apricot", "fish"]
    expected = ["d", "c", "app", "apr", "f"]
    assert shortest_prefix(input) == expected


if __name__ == "__main__":
    print(shortest_prefix(["dog", "cat", "apple", "apricot", "fish"]))
