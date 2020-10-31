import sys


def eprint(msg: str):
    print(msg, file=sys.stderr)


def perfect_number(n: int) -> int:
    if n <= 0 or n > 59271:
        raise ValueError("valid range for n is [1, 59271] got {}".format(n))

    i = 0
    current = 19
    nextPower = 100
    while True:
        if is_perfect_fast(current):
            i += 1

        if i == n:
            return current

        if current == nextPower:
            eprint("passed {}".format(nextPower))
            nextPower *= 10

        current += 1


def is_perfect_fast(n: int) -> bool:
    sum = 0
    while True:
        remaining_places = n // 10
        if remaining_places == 0:
            sum += n
            return sum == 10

        sum += n - remaining_places * 10
        n = remaining_places


if __name__ == "__main__":
    eprint("{}".format(perfect_number(59271)))
