from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        max_area = 0
        max_height = max(height)
        j = len(height) - 1
        i = 0
        counter = 0
        while counter < len(height):
            if max_area < min(height[i], height[j])*(j - i):
                max_area=min(height[i], height[j])*(j - i)
            counter += 1
            if height[i] <= height[j]:
                i += 1
            elif height[j] < height[i]:
                j -= 1
        return max_area



if __name__ == "__main__":
    sol = Solution()
    print(sol.maxArea([1,3,2,5,25,24,5]))

