import math, sympy
from sympy.solvers.diophantine.diophantine import base_solution_linear
from sympy.abc import t

data = "41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,x,x,863,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"
# data = "17,x,13,19"   # 3417
# data = "67,7,x,59,61" # 1261476

def to_buses(data):
    buses = []
    for i, v in enumerate(data.split(',')):
        if v != 'x':
            period = int(v)
            since_depart = 0 if i ==0 else period - i
            since_depart = since_depart % period
            buses.append([since_depart, period])
    return buses

def lcm(a, b):
    return abs(a*b) // math.gcd(a, b)

def is_valid(n, buses):
    for since_depart, period in buses:
        if n % period != since_depart:
            return False
    return True

def get_next_offset_and_period(cur_offset, cur_period, req_offset, req_period):
    # cur_offset + k*cur_period = req_offset + n*req_period
    print("\nsolving for added", req_offset, req_period)
    sol_k, sol_n = base_solution_linear(req_offset - cur_offset, cur_period, -1*req_period, t)
    print("k=", sol_k, "n=", sol_n)
    print("finding new offset for", req_offset, req_period)
    # finding minimal integer t which satisfies a*t + b >= 0
    b, at = sol_k.args
    if at.args:
        a, _ = at.args
        print("a, b", a, b)
        min_t = int(math.ceil(-1 * b / a))
    else:
        min_t = 0
    print("min_t", min_t)
    k = int(round(sol_k.evalf(subs={t:min_t})))
    print("calculating new period", req_offset, req_period)
    return cur_offset + k*cur_period, lcm(cur_period, req_period)

buses = to_buses(data)
cur_offset = 0
cur_period = 1

cur_buses = [] # debug
for req_offset, req_period in buses:
    print("Offset and Period:", cur_offset, cur_period)
    cur_offset, cur_period = get_next_offset_and_period(cur_offset, cur_period, req_offset, req_period)

    cur_buses.append((req_offset, req_period))
    print("Valid:", is_valid(cur_offset, cur_buses))   # I already had this nice function, so why not debug more


print("\nFinal offset:", cur_offset)
