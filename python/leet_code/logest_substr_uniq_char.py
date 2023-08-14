def lengthOfLongestSubstring(s: str) -> int:
    d=set()
    s_l=0
    i = 0
    j = 0
    while j < (len(s)):
        if s[j] not in d:
            d.add(s[j])
            j += 1
            s_l = s_l if s_l > (j - i) else (j - i)
        else:
            d.remove(s[i])
            i += 1
    return len(d) if s_l == 0 else s_l

if __name__ == "__main__":
    print(lengthOfLongestSubstring("aab"))