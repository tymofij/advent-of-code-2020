def get_seat_id(s):
    row = int(s[:7].replace('F','0').replace('B', '1'), 2)
    col = int(s[7:].replace('L','0').replace('R', '1'), 2)
    return row * 8 + col

seat_ids = sorted(get_seat_id(s) for s in open('input.txt').readlines())
print('Max:', max(seat_ids))
for i, v in enumerate(seat_ids):
  if i > 0 and seat_ids[i-1]+1 != v:
    print("Missing:", v-1)



