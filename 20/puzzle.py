from pprint import pprint
import sys

def to_struct(top, left, right, bottom):
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

def flips(a, b, c, d):
    yield from rotations(a, b, c, d) # norm
    yield from rotations(d, b[::-1], c[::-1], a) # top with bottom
    yield from rotations(a[::-1], c, b, d[::-1]) # left and right

tiles = {}
for data in open("input.txt").read().split('\n\n'):
    lines = data.split('\n')
    n = int(lines[0][-5:-1])
    data = lines[1:]
    top = data[0]
    bottom = data[-1]
    left = ''.join([line[0] for line in data])
    right = ''.join([line[-1] for line in data])
    tiles[n] = [to_struct(*variant) for variant in flips(top, left, right, bottom)]


def stitch(left_tile, top_tile):
    # finds tile which fits to the right of this one
    for n in tiles.keys():
        if n in used_tiles:
            continue
        for v in range(12):
            fit = True
            if left_tile and tiles[n][v]['left'] != left_tile['right']:
                fit = False
            if top_tile and tiles[n][v]['top'] != top_tile['bottom']:
                fit = False
            if fit:
                used_tiles.add(n)
                return n, v
    return None

def get_tile(n_v):
    n, v = n_v
    return tiles[n][v]


max_done = 0
for n in tiles.keys():
    for v in range(12):
        print(n, v)
        top_left = n, v
        used_tiles = set({n})
        matrix = [[top_left]] # matrix of 12x12, each element a tuple of (tile_n, variant_n)
        done = False
        for j in range(1, 12):
            next_tile_nv = stitch(get_tile(matrix[0][j-1]), None)
            if next_tile_nv:
                matrix[0].append(next_tile_nv)
            else:
                done = True
                break
        if done:
            continue # we could not even finish the first row
        for i in range(1, 12):
            if done:
                break
            next_tile_nv = stitch(None, get_tile(matrix[i-1][0]))
            if next_tile_nv:
                matrix.append([next_tile_nv])
            else:
                done = True
                break
            for j in range(1, 12):
                next_tile_nv = stitch(get_tile(matrix[i][j-1]), get_tile(matrix[i-1][j]))
                if next_tile_nv:
                    matrix[i].append(next_tile_nv)
                else:
                    done = True
                    break

        max_done = max(max_done, len(used_tiles))
        print(' ', len(used_tiles))

        if len(used_tiles) == 144:
            print("FUCK YEAH!")
            sys.exit()

pprint(max_done)