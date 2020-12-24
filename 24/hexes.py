DIRS = {
    'e': (2,  0),
    'w': (-2, 0),
    'se': ( 1, -1),
    'sw': (-1, -1),
    'ne': ( 1,  1),
    'nw': (-1,  1),
}

data = [s.strip() for s in open('input.txt').readlines()]

def to_coords(line):
    x, y = 0, 0
    acum = ''
    for char in line:
        acum += char
        if acum in DIRS:
            dx, dy = DIRS[acum]
            x += dx
            y += dy
            acum = ''
    return x, y

def count_black_neighbours(x, y, data):
    res = 0
    for dx, dy in DIRS.values():
        if data.get((x+dx, y+dy)):
            res += 1
    return res

def count_blacks(data):
    return sum([1 for elem in data.values() if elem])

assert to_coords('nwwswee') == (0, 0)

tiles = {}
for line in data:
    coords = to_coords(line)
    if coords not in tiles:
        tiles[coords] = True # black
    else:
        tiles[coords] = not tiles[coords]

print('Initially black tiles: ', count_blacks(tiles))

for day in range(101):
    print(f'Day {day}:', count_blacks(tiles))
    new_data = {}
    min_x = min_y = float('inf')
    max_x = max_y = float('-inf')
    for (x, y), val in tiles.items():
        if val:
            min_x = min(x, min_x)
            min_y = min(y, min_y)
            max_x = max(x, max_x)
            max_y = max(y, max_y)
    for x in range(min_x-2, max_x+2+1):
        for y in range(min_y-1, max_y+2):
            blacks = count_black_neighbours(x, y, tiles)
            if tiles.get((x, y)):
                if not(blacks == 0 or blacks > 2):
                    new_data[x, y] = True
            else:
                if blacks == 2:
                    new_data[x, y] = True
    tiles = new_data

