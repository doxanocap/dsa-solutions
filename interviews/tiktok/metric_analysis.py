class Solution:
    def validSubarraySize(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        stack = []

        left = [i for i in range(n)]
        for i in range(n):
            while stack and nums[stack[-1]] >= nums[i]:
                left[i] = left[stack.pop()]
            stack.append(i)

        right = [i for i in range(n)]
        for i in range(n-1, -1, -1):
            while stack and nums[stack[-1]] >= nums[i]:
                right[i] = right[stack.pop()]
            stack.append(i)

        for i in range(n):
            k = right[i] - left[i] + 1
            if nums[i] > threshold / k:
                return k

        return -1
