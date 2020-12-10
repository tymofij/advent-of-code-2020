numbers = [int(s) for s in  open('i.txt').readlines()]
numbers.append(0)
numbers.sort()

def count_ways(nums):
    res = 0
    if len(nums) <= 2:
        return 1
    if nums[-1] - nums[-2] <= 3:
        res += count_ways(nums[:-1])
    try:
        if nums[-1] - nums[-3] <= 3:
            res += count_ways(nums[:-2])
        if nums[-1] - nums[-4] <= 3:
            res += count_ways(nums[:-3])
    except IndexError:
        pass
    return res

print(count_ways(numbers))