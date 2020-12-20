from pprint import pprint
import math, sys

VARIATIONS = 8
FILENAME = "input.txt"
# FILENAME = "demo.txt"

DEBUG = 0
def log(*args):
    if DEBUG:
        print(*args)


def rotate(matrix, n):
    res = matrix
    for i in range(n):
        res = list(zip(*res[::-1]))
    return [
        ''.join(line) for line in res
    ]


def flip(matrix):
    return [line[::-1] for line in matrix]


def to_struct(matrix):
    top = matrix[0]
    bottom = matrix[-1]
    left = ''.join([line[0] for line in matrix])
    right = ''.join([line[-1] for line in matrix])
    assert top[0] == left[0]
    assert top[-1] == right[0]
    assert bottom[0] == left[-1]
    assert bottom[-1] == right[-1]
    return {
        'top': top,
        'left': left,
        'right': right,
        'bottom': bottom,
        'matrix': matrix
    }


def variants(matrix):
    for i in range(4):
        yield rotate(matrix, i)
    flipped = flip(matrix)
    for i in range(4):
        yield rotate(flipped, i)

tiles = {}
for data in open(FILENAME).read().split('\n\n'):
    lines = data.split('\n')
    n = int(lines[0][-5:-1])
    data = lines[1:]
    tiles[n] = [to_struct(variant) for variant in variants(data)]

SIZE = round(math.sqrt(len(tiles)))
log('prepared variants')

# precalc tile cache
for n in tiles.keys():
    for v in range(VARIATIONS):
        tiles[n][v]['fit_right'] = set()
        tiles[n][v]['fit_below'] = set()
        for nn in tiles.keys():
            if n == nn:
                continue
            for mm in range(VARIATIONS):
                if tiles[n][v]['right'] == tiles[nn][mm]['left']:
                    tiles[n][v]['fit_right'].add((nn, mm))
                if tiles[n][v]['bottom'] == tiles[nn][mm]['top']:
                    tiles[n][v]['fit_below'].add((nn, mm))
log('precalculated cache')


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
    # 'options': [(1951, 6)],  # to force exactly the answer as in demo solution
    'exclude': set(),
    'n': None, # - to be filled in
    'v': None, # - to be filled in
}]


max_len = 0
min_roots = len(chain[0]['options'])
while chain:
    if DEBUG:
        if len(chain) > max_len:
            max_len = len(chain)
            print(f'longer chain: roots {min_roots}, max chain {max_len}')
        if len(chain[0]['options']) < min_roots:
            min_roots = len(chain[0]['options'])
            print(f'fewer roots:  roots {min_roots}, max chain {max_len}')

    log('chain', chain)
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


if DEBUG:
    print('------')
    print(f'max length {max_len} out of {len(chain)}')
    for link in chain:
        print(link['n'], link['v'], link['exclude'])

if chain:
    print('!!', chain[0]['n'] * chain[-1]['n'] * chain[SIZE-1]['n'] * chain[SIZE*(SIZE-1)]['n'])