data = "41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,x,x,863,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"

# data = "17,x,13,19"
buses = {}
for k, v in enumerate(data.split(',')):
    if v != 'x':
        buses[k] = int(v)

print(buses)

def is_valid(n, buses):
    for k, v in buses.items():
        since_depart = n % v
        to_next = v - since_depart if since_depart else 0
        if to_next != k:
            return False
    return True

print(is_valid(3417, buses))
print(is_valid(10010, buses))