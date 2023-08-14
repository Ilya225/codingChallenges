import sys

def maxMin(k, arr):
    arr.sort()

    minim = 10**9
    for i in range(0, len(arr) - k):
        if minim > arr[k + i - 1] - arr[i]:
            minim = arr[k + i - 1] - arr[i]

    return minim





if __name__ == "__main__":

    n = int(input())

    k = int(input())

    arr = []

    for _ in range(n):
        arr_item = int(input())
        arr.append(arr_item)

    result = maxMin(k, arr)

    print(result)