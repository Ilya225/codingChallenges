def luckBalance(k, contests):

    bottomK = sorted(map(lambda c: c[0], filter(lambda c: c[1] == 1, contests)))
    bottomK = bottomK[:len(bottomK) - k]

    result = sum(map(lambda c: c[0], contests)) - sum(bottomK)*2

    return result



if __name__ == '__main__':

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    contests = []

    for _ in range(n):
        contests.append(list(map(int, input().rstrip().split())))

    result = luckBalance(k, contests)

    print(result)
