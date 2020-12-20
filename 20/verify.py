""" I messed up matrix rotation logic, this is debug to find what went wrong
"""

from puzzle import *

sol = []
for triplet in open('sol.txt').read().split('\n\n'):
    one, two, three = [], [], []
    for line in triplet.split('\n'):
        a, b, c = line.split()
        one.append(a)
        two.append(b)
        three.append(c)
    sol.extend([one, two, three])

i = 0
for data in sol:
    i +=1
    for line in data:
        print(line)
    found = False
    for n in tiles.keys():
        for v in range(VARIATIONS):
            t = tiles[n][v]
            if t['matrix'] == data:
                print(f'demo solution tile {i} fits variant ({n}, {v})')
                found = True
    if not found:
        print('unmatched tile!')
    print()


print('total tiles matched:', i)
