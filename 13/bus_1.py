start = 1001938
data = "41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,x,x,863,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"
buses = [int(s) for s in data.split(',') if s != 'x']

min_delay = start
min_bus = None
for bus in buses:
    delay = bus - (start % bus)
    if delay < min_delay:
        min_delay = delay
        min_bus = bus

print(min_bus * min_delay)