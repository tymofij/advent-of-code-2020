import math

data = [(s[0], int(s[1:])) for s in open('input.txt').readlines()]

x = 0
y = 0
dx = 10
dy = 1

angle = 0  # in degrees
for rule, n in data:
    if rule == 'N':
        dy += n
    elif rule == 'S':
        dy -= n
    elif rule == 'E':
        dx += n
    elif rule == 'W':
        dx -= n
    elif rule == 'F':
        x += dx * n
        y += dy * n
    elif rule in ('L', 'R'):
        d_rad = rad = (math.pi * n) / 180
        radius = math.sqrt(dx**2 + dy**2)
        if rule == 'R':
            d_rad *= -1
        rad = math.atan2(dy, dx) + d_rad
        dy = radius * math.sin(rad)
        dx = radius * math.cos(rad)

print(round(abs(x) + abs(y)))

