data = open('input.txt').read()

# Sum of Number of unique answers in each group
groups = [s.replace('\n', '') for s in data.split('\n\n')]
print(sum(len(set(g)) for g in groups))

def intersect(*args):
    res = set(args[0])
    for a in args:
        res = res & set(a)
    return len(res)

# Sum of Number of intersecting answers in each group
groups = [s.split() for s in data.split('\n\n')]
print(sum([intersect(*g) for g in groups]))
