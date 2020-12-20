from puzzle import *

def wo_borders(matrix):
    return [line[1:-1] for line in matrix[1:-1]]

PIECE_DIM = len(tiles[chain[0]['n']][0]['top']) - 2
PHOTO_SIZE = SIZE * PIECE_DIM

photo = []
for i in range(PHOTO_SIZE):
    photo.append(['']* PHOTO_SIZE)

for link_i in range(SIZE):
    for link_j in range(SIZE):
        link = chain[link_i*SIZE + link_j]
        matrix = wo_borders(tiles[link['n']][link['v']]['matrix'])
        for i in range(PIECE_DIM):
            for j in range(PIECE_DIM):
                photo[link_i*PIECE_DIM + i][link_j*PIECE_DIM + j] = matrix[i][j]

MONSTER = [
'                  # ',
'#    ##    ##    ###',
' #  #  #  #  #  #   '
]
MONSTER_W = len(MONSTER[0])
MONSTER_H = len(MONSTER)


def get_monster_coords(photo):
    res = []
    for i in range(PHOTO_SIZE-MONSTER_H+1):
        for j in range(PHOTO_SIZE-MONSTER_W+1):
            found = True
            for mi in range(MONSTER_H):
                for mj in range(MONSTER_W):
                    # print(i, j, mi, mj)
                    if MONSTER[mi][mj] == '#' and photo[i+mi][j+mj] != '#':
                        found = False
            if found:
                res.append([i, j])
    return res


def reveal_monster(photo, coord_pairs):
    photo = [list(s) for s in photo]
    for i, j in coord_pairs:
        for mi in range(MONSTER_H):
            for mj in range(MONSTER_W):
                if MONSTER[mi][mj] == '#':
                    photo[i+mi][j+mj] = 'O'
    return photo


for photo_var in variants(photo):
    monster_coords = get_monster_coords(photo_var)
    if monster_coords:
        photo_w_monsters = reveal_monster(photo_var, monster_coords)

        s = 0
        for i in range(PHOTO_SIZE):
            for j in range(PHOTO_SIZE):
                if photo_w_monsters[i][j] == '#':
                    s+=1
        print('Sea roughness:', s)


