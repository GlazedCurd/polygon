#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import random
import sys

def random_array(size, st, end):
    random_array = [random.randint(st, end) for _ in range(size)]
    return random_array

def main():
    # passing seed through command line arguments
    random.seed(int(sys.argv[1]))
    array_size = random.randint(5, 15)
    query = random.randint(-12, 12)
    array = random_array(array_size, -10, 10)
    array.sort()
    print(array_size)
    print(" ".join(str(i) for i in array))
    print(query)

if __name__ == "__main__":
    main()