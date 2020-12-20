from pprint import pprint
import math, sys

VARIATIONS = 12

def to_struct(top, right, bottom, left):
    return {
        'top': top,
        'left': left,
        'right': right,
        'bottom': bottom,
    }

def rotations(a, b, c, d):
    yield a, b, c, d
    yield b, c, d, a
    yield c, d, a, b
    yield d, a, b, c

def flips(top, right, bottom, left):
    yield from rotations(top, right, bottom, left) # norm
    # vertical: replace top with bottom
    yield from rotations(bottom, right[::-1], top, left[::-1])
    # horizontal: replace left with right
    yield from rotations(top[::-1], left, bottom[::-1], right)

tiles = {}
for data in open("demo.txt").read().split('\n\n'):
    lines = data.split('\n')
    n = int(lines[0][-5:-1])
    data = lines[1:]
    top = data[0]
    bottom = data[-1]
    left = ''.join([line[0] for line in data])
    right = ''.join([line[-1] for line in data])
    tiles[n] = [to_struct(*variant) for variant in flips(top, right, bottom, left)]

SIZE = round(math.sqrt(len(tiles)))

def stitch(left_tile, top_tile, exclude_tiles):
    # finds tile which fits to the right of this one
    res = []
    for n in tiles.keys():
        if n in exclude_tiles:
            continue
        for v in range(VARIATIONS): # variations
            fit = True
            if left_tile and tiles[n][v]['left'] != left_tile['right']:
                fit = False
            if top_tile and tiles[n][v]['top'] != top_tile['bottom']:
                fit = False
            if fit:
                res.append([n, v])
    return res

def get_tile(link):
    if link is None:
        return None
    return tiles[link['n']][link['v']]


chain = [{
    'options': stitch(None, None, {}),
    'exclude': set(),
    'n': None, # - to be filled in
    'v': None, # - to be filled in
}]

# m = 0
# while chain:
#     print(len(chain))
#     print('chain', chain)
#     m = max(m, len(chain))
#     if len(chain) == SIZE*SIZE:
#         break # we're done
#     i = len(chain)  # next elem to assign
#     cur = chain[i-1]
#     if not cur['options']:
#         print('out of options, taking step back')
#         chain.pop() # step back
#         continue
#     n, v = cur['options'].pop()
#     cur['n'] = n
#     cur['v'] = v
#     cur['exclude'] = {n} if i == 1 else chain[i-2]['exclude'] | {n}
#     print('evaluating', cur)
#     left = chain[i-1] if i % SIZE != 0 else None
#     top = chain[i-SIZE] if i >= SIZE else None
#     next_ones = stitch(get_tile(left), get_tile(top), cur['exclude'])
#     if next_ones:
#         print('got options, extending')
#         chain.append({
#             'options': next_ones,
#             'exclude': None,
#             'n': None, # - to be filled in
#             'v': None, # - to be filled in
#         })
#     else:
#         print('could not find next tile')

# print('------')
# print(m, len(chain))
# for link in chain:
#     print(link['n'], link['v'], link['exclude'])

sol = []
for triplet in open('sol.txt').read().split('\n\n'):
    one, two, three = [], [], []
    for line in triplet.split('\n'):
        a, b, c = line.split()
        one.append(a)
        two.append(b)
        three.append(c)
    sol.extend([one, two, three])

for data in sol:
    top = data[0]
    bottom = data[-1]
    left = ''.join([line[0] for line in data])
    right = ''.join([line[-1] for line in data])
    found = False
    for n in tiles.keys():
        for v in range(VARIATIONS):
            t = tiles[n][v]
            if t['left'] == left and t['right'] == right and t['top'] == top and t['bottom'] == bottom:
                print('FOUND!')
                found = True
    if not found:
        for line in data:
            print(line)



