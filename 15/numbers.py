data = (10,16,6,0,1,17)

last_spoken = {v: k+1 for k, v in enumerate(data[:-1])}

prev = data[-1]
turn = len(data)

while turn < 30000000:
    if prev not in last_spoken:
        new = 0
    else:
        new = turn - last_spoken[prev]

    last_spoken[prev] = turn

    prev = new
    turn += 1

print(turn, prev)