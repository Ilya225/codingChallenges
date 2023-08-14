def string_permutation(perm_str, remainder):
    if remainder == "":
        print(perm_str)

    for i in range(len(remainder)):
        string_permutation(perm_str + remainder[i], remainder[0:i] + remainder[i+1:])


if __name__ == '__main__':
    string_permutation("", "helo")