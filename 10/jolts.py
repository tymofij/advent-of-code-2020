numbers = [int(s) for s in open('input.txt').readlines()]
numbers.append(0)
numbers.sort()

cache = {}

def count_ways(nums):
    last = nums[-1]
    if last in cache:
        return cache[last]
    res = 0
    if len(nums) <= 2:
        return 1
    if last - nums[-2] <= 3:
        res += count_ways(nums[:-1])
    if len(nums) >=3 and last - nums[-3] <= 3:
        res += count_ways(nums[:-2])
    if len(nums) >=4 and last - nums[-4] <= 3:
        res += count_ways(nums[:-3])
    cache[last] = res
    return res

print(count_ways(numbers))