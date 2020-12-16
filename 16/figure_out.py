# from data import rules, your_ticket, nearby_tickets
from data import rules, your_ticket, nearby_tickets
from collections import defaultdict

"""
I'm not proud of this code
"""


def is_valid(ticket):
    for n in ticket:
        valid = False
        for a, b in rules.values():
            if a[0] <= n <= a[1] or b[0] <= n <= b[1]:
                valid = True
        if not valid:
            return False
    return True

valid_tickets = [t for t in nearby_tickets if is_valid(t)]


def rule_applies(name, ranges, pos, tickets):
    a, b = ranges
    for t in tickets:
        valid = False
        if a[0] <= t[pos] <= a[1] or b[0] <= t[pos] <= b[1]:
            valid = True
        if not valid:
            return False
    return True

n_fields = len(valid_tickets[0])

rule_positions = defaultdict(set)

for name, ranges in rules.items():
    for i in range(n_fields):
        if rule_applies(name, ranges, i, valid_tickets):
            rule_positions[name].add(i)


unclear = set(rules.keys())
while unclear:
    print(unclear)
    print('rule_positions', rule_positions)
    print()
    clear_name, clear_position = [
        (name, positions) for (name, positions) in rule_positions.items()
        if len(positions) == 1 and (name in unclear)
    ][0]
    clear_position = list(clear_position)[0]
    print('clear_name', clear_name)
    print('clear_position', clear_position)
    print()
    unclear.remove(clear_name)
    for name in rule_positions.keys():
        if name != clear_name and clear_position in rule_positions[name]:
            print(f'removing {clear_position} from {name}')
            rule_positions[name].remove(clear_position)

departure_positions = [list(positions)[0] for (name, positions) in rule_positions.items() if 'departure' in name]

print('Departure positions', departure_positions)

s = 1
for i in departure_positions:
    s *= your_ticket[i]

print('Departure values multiplied:', s)
