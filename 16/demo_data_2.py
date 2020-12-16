rules = {x.split(':')[0] : x.split(':')[1].strip() for x in """class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19""".split("\n") }

def to_dep(s):
    return int(s.split('-')[0]), int(s.split('-')[1])

for k in rules:
    rules[k] = [to_dep(x) for x in rules[k].split(' or ')]

def to_ints(s):
    return [int(x) for x in s.split(',')]

your_ticket = to_ints("11,12,13")

nearby_tickets = [to_ints(x) for x in """3,9,18
15,1,5
5,14,9""".split("\n")]

