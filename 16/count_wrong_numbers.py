# from data import rules, your_ticket, nearby_tickets
from data import rules, your_ticket, nearby_tickets


def count_invalid_vals(ticket):
    res = 0
    for n in ticket:
        valid = False
        for a, b in rules.values():
            if a[0] <= n <= a[1] or b[0] <= n <= b[1]:
                valid = True
        if not valid:
            res += n
    return res

sum_invalid_numbers = sum([count_invalid_vals(ticket) for ticket in nearby_tickets])

print(sum_invalid_numbers)
