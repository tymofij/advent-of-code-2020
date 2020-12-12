import math

data = [(s[0], int(s[1:])) for s in open('input.txt').readlines()]

x = 0
y = 0
angle = 0
for rule, n in data:
    if rule == 'N':
        y += n
    elif rule == 'S':
        y -= n
    elif rule == 'E':
        x += n
    elif rule == 'W':
        x -= n
    elif rule == 'F':
        angle += 360
        angle = angle % 360
        rad = (math.pi * angle) / 180
        x += n * math.cos(rad)
        y += n * math.sin(rad)
    elif rule == 'L':
        angle += n
    elif rule == 'R':
        angle -= n

print(round(abs(x) + abs(y)))

