from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        maxHeight = 0
        tmpCapacity = 0
        fullCapacity = 0

        for h in height:
            if h >= maxHeight:
                maxHeight = h
                fullCapacity += tmpCapacity
                tmpCapacity = 0
            tmpCapacity += (maxHeight - h)

        rightMaxHeight = 0
        tmpCapacity = 0
        # rightIndex = len(height) - 1

        for rightIndex in range(len(height) -1, -1, -1):
            if height[rightIndex] >= rightMaxHeight:
                rightMaxHeight = height[rightIndex]
                fullCapacity += tmpCapacity
                tmpCapacity = 0
            tmpCapacity += (rightMaxHeight - height[rightIndex])
            if height[rightIndex] == maxHeight:
                break

        return fullCapacity

    def trapOn(self, height: List[int]) -> int:
        if not len(height):
            return 0


        left = 0
        right = len(height) - 1

        leftMax = height[left]
        rightMax = height[right]
        capacity = 0

        while left != right:
            if height[left] > height[right]:
                right -= 1
                if height[right] > rightMax:
                    rightMax = height[right]
                capacity += rightMax - height[right]
            else:
                left += 1
                if height[left] > leftMax:
                    leftMax = height[left]
                capacity+= leftMax - height[left]

        return capacity


if __name__ == '__main__':
    sol = Solution()
    print(sol.trapOn([4,2,3]))