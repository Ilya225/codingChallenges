from typing import List


class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        words_set = set(words)
        word_check = set()

        subStr = ''
        for c in s:
            subStr += c
            if subStr in words_set:
                word_check.add(subStr)
                subStr = ''
            else:
                pass


if __name__ == '__main__':
    sol = Solution()