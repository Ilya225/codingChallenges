import math
import os
import random
import re
import sys

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries):
    array = [0 for i in range(0,n)]

    for q in queries:
        array[q[0]-1] += q[2]
        if q[1] < len(array):
            array[q[1]] -= q[2]

    max = 0
    current = 0
    for a in array:
        current += a
        if current > max:
            max = current

    return max

if __name__ == '__main__':
    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    queries = []

    for _ in range(m):
        queries.append(list(map(int, input().rstrip().split())))

    result = arrayManipulation(n, queries)

    print(result)

