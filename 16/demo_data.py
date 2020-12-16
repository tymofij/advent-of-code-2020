rules = {x.split(':')[0] : x.split(':')[1].strip() for x in """class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50""".split("\n") }

def to_dep(s):
    return int(s.split('-')[0]), int(s.split('-')[1])

for k in rules:
    rules[k] = [to_dep(x) for x in rules[k].split(' or ')]

def to_ints(s):
    return [int(x) for x in s.split(',')]

your_ticket = to_ints("7,1,14")

nearby_tickets = [to_ints(x) for x in """7,3,47
40,4,50
55,2,20
38,6,12""".split("\n")]

