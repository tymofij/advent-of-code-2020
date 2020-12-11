from copy import deepcopy
seats = [list(s.strip()) for s in open('input.txt').readlines()]

FLOOR = '.'
EMPTY = 'L'
TAKEN = '#'

def count_visible(seats, i, j, visibility_limit=None):
    res = 0
    directions = [
        (-1, -1), (-1, 0), (-1, 1),
        (0,  -1),          (0,  1),
        (1,  -1), (1,  0), (1 , 1)
    ]
    for (di, dj) in directions:
        ni, nj = i+di, j+dj
        steps = 1
        while ((0 <= ni <= len(seats)-1) and
               (0 <= nj <= len(seats[0])-1)):
            if seats[ni][nj] == EMPTY:
                break
            if seats[ni][nj] == TAKEN:
                res += 1
                break
            if visibility_limit and steps >= visibility_limit:
                break
            ni, nj = ni+di, nj+dj
            steps += 1
    return res


def next_state(seats, tolerance, visibility_limit=None):
    new = deepcopy(seats)
    for i in range(len(seats)):
        for j in range(len(seats[0])):
            if seats[i][j] == FLOOR:
                continue
            neighbours = count_visible(seats, i, j, visibility_limit)
            if seats[i][j] == EMPTY and neighbours == 0:
                new[i][j] = TAKEN
            if seats[i][j] == TAKEN and neighbours >= tolerance:
                new[i][j] = EMPTY
    return new


def show(seats):
    for line in seats:
        print(''.join(line))
    print()


def occupied_seats_when_stabilized(seats, tolerance, visibility_limit=None):
    prev_seats = seats
    seats = next_state(seats, tolerance, visibility_limit)
    while seats != prev_seats:
        # show(seats)
        prev_seats = seats
        seats = next_state(seats, tolerance, visibility_limit)

    res = 0
    for line in seats:
        for seat in line:
            if seat == TAKEN:
                res += 1
    return res

print("part 1:", occupied_seats_when_stabilized(seats, 4, 1))
print("part 2:", occupied_seats_when_stabilized(seats, 5))
