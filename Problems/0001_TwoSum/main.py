
# The correc Solution.

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        pair_idx = {}
        for i, num in enumerate(nums):
            if target - num in pair_idx:
                return [i, pair_idx[target - num]]
            pair_idx[num] = i


numbers = [2,7,11,15,5]
target = 7

solution = Solution()
result = solution.twoSum(numbers , target) 
print(result)

