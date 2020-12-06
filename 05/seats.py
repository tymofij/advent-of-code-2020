def get_seat_id(s):
  return int(s.replace('F','0').replace('B','1')
              .replace('L','0').replace('R','1'), 2)

seat_ids = sorted(get_seat_id(s) for s in open('input.txt').readlines())
print('Max:', seat_ids[-1])

for i, v in enumerate(seat_ids):
  if i > 0 and seat_ids[i-1]+1 != v:
    print("Missing:", v-1)



