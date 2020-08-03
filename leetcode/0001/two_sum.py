# coding=utf-8

class Solution:
    def twoSum(self, nums, target):
        m = {}
        for i in range(len(nums)):
            if m.get(target - nums[i]) is not None:
                return [m.get(target - nums[i]), i]

            m[nums[i]] = i
        return []


if __name__ == '__main__':
    print(Solution().twoSum([2, 7, 11, 2], 9))
