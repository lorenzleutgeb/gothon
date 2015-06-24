# The Computer Language Benchmarks Game
# http://benchmarksgame.alioth.debian.org/
#
# contributed Daniel Nanz

import sys
import gmpy2


def get_pistring(n, pi=gmpy2.pi):

    return str(pi(int(3.35 * n))).replace('.', '')[0 : n]


def main(n, width=10, line='{}\t:{}'):

    pistring = get_pistring(n)
    for i in range(0, n - width + 1, width):
        print(line.format(pistring[i : i + width], i + width))
    if n % width > 0:
        print(line.format(pistring[-(n % width) : ].ljust(width), n))


main(int(sys.argv[1]))
