from pprint import pprint
import math, sys

VARIATIONS = 12

def to_struct(top, right, bottom, left):
    assert top[0] == left[0]
    assert top[-1] == right[0]
    assert bottom[0] == left[-1]
    assert bottom[-1] == right[-1]
    return {
        'top': top,
        'left': left,
        'right': right,
        'bottom': bottom,
    }

def rotations(top, right, bottom, left):
    yield top, right, bottom, left
    top, right, bottom, left = right, bottom[::-1], left, top[::-1]
    yield top, right, bottom, left
    top, right, bottom, left = right, bottom[::-1], left, top[::-1]
    yield top, right, bottom, left
    top, right, bottom, left = right, bottom[::-1], left, top[::-1]
    yield top, right, bottom, left

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

print('prepared variants')
SIZE = round(math.sqrt(len(tiles)))


# precalc tile cache
for n in tiles.keys():
    for v in range(VARIATIONS):
        tiles[n][v]['fit_right'] = set()
        tiles[n][v]['fit_below'] = set()
        for nn in tiles.keys():
            if n == nn:
                continue
            for mm in range(VARIATIONS):
                if tiles[n][v]['left'] == tiles[nn][mm]['right']:
                    tiles[n][v]['fit_right'].add((nn, mm))
                if tiles[n][v]['bottom'] == tiles[nn][mm]['top']:
                    tiles[n][v]['fit_below'].add((nn, mm))
print('precalculated cache')

def stitch(left_tile, top_tile, exclude_tiles):
    # finds tile which fits to the right of this one
    if left_tile and not top_tile:
        res = left_tile['fit_right']
    elif top_tile and not left_tile:
        res = top_tile['fit_below']
    elif top_tile and left_tile:
        fit_left = left_tile['fit_right']
        fit_below = top_tile['fit_below']
        res = fit_left & fit_below
    elif not top_tile and not left_tile:
        res = []
        for n in tiles.keys():
            for v in range(VARIATIONS):
                res.append((n, v))
    return [(n, v) for (n, v) in list(res) if n not in exclude_tiles]

def get_tile(link):
    if link is None:
        return None
    return tiles[link['n']][link['v']]


chain = [{
    'options': stitch(None, None, set()),
    'exclude': set(),
    'n': None, # - to be filled in
    'v': None, # - to be filled in
}]

DEBUG = 0
def log(*args):
    if DEBUG:
        print(*args)

m = 0
while chain:
    log(len(chain))
    log('chain', chain)
    m = max(m, len(chain))
    i = len(chain)  # next elem to assign
    cur = chain[i-1]
    if not cur['options']:
        log('out of options, taking step back')
        chain.pop() # step back
        continue
    n, v = cur['options'].pop()
    cur['n'] = n
    cur['v'] = v
    cur['exclude'] = {n} if i == 1 else chain[i-2]['exclude'] | {n}
    if len(chain) == SIZE*SIZE:
        break # we're done
    log('evaluating', cur)
    left = chain[i-1] if i % SIZE != 0 else None
    top = chain[i-SIZE] if i >= SIZE else None
    next_ones = stitch(get_tile(left), get_tile(top), cur['exclude'])
    if next_ones:
        log('got options, extending')
        chain.append({
            'options': next_ones,
            'exclude': None,
            'n': None, # - to be filled in
            'v': None, # - to be filled in
        })
    else:
        log('could not find next tile')

print('------')
print(f'max length {m} out of {len(chain)}')
for link in chain:
    print(link['n'], link['v'], link['exclude'])

print('!!', chain[0]['n'] * chain[-1]['n'] * chain[SIZE-1]['n'] * chain[SIZE*(SIZE-1)]['n'])