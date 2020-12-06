data = open('input.txt').read()
groups = [[set(_) for _ in s.split()] for s in data.split('\n\n')]

# Sum of Number of unique answers in each group
print(sum([len(set.union(*g)) for g in groups]))

# Sum of Number of intersecting answers in each group
print(sum([len(set.intersection(*g)) for g in groups]))