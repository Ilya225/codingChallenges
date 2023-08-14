from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:

        i = 0
        jumps = 0
        while i < len(nums)-1:
            max_index = 0
            for j in range(1, nums[i]+1):
                if i + j < len(nums):
                    if max_index == 0:
                        max_index = i+j
                    elif nums[i + j] >= nums[max_index]:
                        max_index = i + j
                else:
                    jumps += 1
                    return jumps
            i = max_index
            jumps += 1

        return jumps

if __name__ == '__main__':
    sol = Solution()
    print(sol.jump([2,3,1]))