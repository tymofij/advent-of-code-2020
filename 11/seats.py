from copy import deepcopy
seats = [list(s.strip()) for s in open('input.txt').readlines()]

FLOOR = '.'
EMPTY = 'L'
TAKEN = '#'


def count_taken_around(seats, i, j):
    res = 0
    for i_step in (-1, 0, 1):
        for j_step in (-1, 0, 1):
            if (0 <= i+i_step <= len(seats)-1 and
               0 <= j+j_step <= len(seats[0])-1 and
               not (i_step==0 and j_step==0) and
               seats[i+i_step][j+j_step] == TAKEN
               ):
                res += 1
    return res


def next_state(seats):
    new = deepcopy(seats)
    for i in range(len(seats)):
        for j in range(len(seats[0])):
            if seats[i][j] == FLOOR:
                continue
            neighbours = count_taken_around(seats, i, j)
            if seats[i][j] == EMPTY and neighbours == 0:
                new[i][j] = TAKEN
            if seats[i][j] == TAKEN and neighbours >= 4:
                new[i][j] = EMPTY
    return new

def show(seats):
    for line in seats:
        print(''.join(line))
    print()

prev_seats = deepcopy(seats)
seats = next_state(seats)
while seats != prev_seats:
    prev_seats = deepcopy(seats)
    seats = next_state(seats)

res = 0
for line in seats:
    for seat in line:
        if seat == TAKEN:
            res += 1
print(res)
