from collections import defaultdict
import re

color_contents = defaultdict(list)
color_ancestors = defaultdict(list)

for line in open('input.txt').readlines():
    color, rest = line.split(' bags contain ')
    if rest == "no other bags.":
        continue
    bag_amounts = re.findall('(\d+) ([^\.,]+) bags?', rest)
    color_contents[color] = bag_amounts
    for _, name in bag_amounts:
        color_ancestors[name].append(color)

# in how many different bags we can put a shiny gold bag
stack = ['shiny gold']
ancestors = set()
while stack:
    color = stack.pop(0)
    for ancestor in color_ancestors[color]:
        ancestors.add(ancestor)
        stack.append(ancestor)

print(len(set(ancestors)))

# how many other bags fit in a shiny gold bag
def count_bags_inside(color):
    contents = color_contents[color]
    return 1 + sum([int(n)*count_bags_inside(color) for (n, color) in contents])

print(count_bags_inside('shiny gold')-1)
