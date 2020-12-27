import operator
from functools import reduce
from collections import defaultdict, Counter

lines = [s.strip().strip(")").split("(contains ") for s in open('input.txt').readlines()]

# allergen => list of occurences
notices = defaultdict(list)
counter = Counter()

for ingredients, allergens in lines:
    ingredients = set(ingredients.split())
    counter.update(ingredients)
    allergens = allergens.split(", ")
    for allergen in allergens:
        notices[allergen].append(ingredients)

for allergen, sightings in notices.items():
    notices[allergen] = reduce(operator.and_, sightings)

# ingredient => allergen
solved = {}

while len(solved) != len(notices):
    for allergen, sightings in notices.items():
        if allergen in solved:
            continue
        notices[allergen] -= set(solved.keys())
        if len(notices[allergen]) == 1:
            food = notices[allergen].pop()
            solved[food] = allergen

print("Occurences of safe ingredients:", sum([v for (k, v) in counter.items() if k not in solved]))
print("Canonical dangerous ingredient list:",
    ",".join([food for (food, allergen) in  sorted(solved.items(), key=lambda e: e[1])])
)
