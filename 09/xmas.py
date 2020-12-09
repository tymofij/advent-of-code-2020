PREFIX = 25  # brute force FTW

numbers = [int(s) for s in  open('input.txt').readlines()]

def is_valid(nums, n):
    for i in nums:
        for j in nums:
            if i+j == n:
                return True
    return False

for i in range(PREFIX, len(numbers)):
    if not is_valid(numbers[max(i-PREFIX, 0): i], numbers[i]):
        X = numbers[i]
        break

for i in range(len(numbers)):
    for j in range(i+2, len(numbers)):
        if sum(numbers[i:j]) == X:
            r = numbers[i:j]
            print(min(r) + max(r))

