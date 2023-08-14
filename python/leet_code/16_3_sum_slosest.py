from typing import List


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        diff = None
        sum = None

        nums = sorted(nums)

        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                d = nums[i] + nums[j]

                # binary search
                low = j + 1
                high = len(nums) - 1
                middle = low + (high - low) // 2

                while low <= high:
                    new_sum = nums[middle] + d
                    if diff is None:
                        diff = abs(new_sum - target)
                        sum = new_sum
                    else:
                        new_diff = min(abs(new_sum - target), diff)
                        if new_diff < diff:
                            sum = new_sum
                            diff = abs(new_sum - target)

                    if new_sum < target:
                        low = middle + 1
                    else:
                        high = middle - 1

                    middle = low + (high - low) // 2

        return sum


if __name__ == '__main__':
    sol = Solution()
    # print(sol.threeSumClosest([4, 0, 5, -5, 3, 3, 0, -4, -5], -2))
    print(sol.threeSumClosest([0,3,97,102,200], 300))
