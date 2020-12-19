
to_update = []

rules = {}
for s in open('rules.txt').readlines():
    n, text = s.split(':')
    n = int(n)
    if '"' in text:
        rules[n] = {'regexp': text.strip().strip('"')}
        to_update.append(n)
    else:
        rules[n] = {'matches': []}
        for match in text.strip().split('|'):
            rules[n]['matches'].append([int(_) for _ in match.split()])

while to_update:
    solved_n = to_update.pop()
    for n, rule in rules.items():
        if rule.get('regexp'):
            continue
        for m in rule['matches']:
            for i in range(len(m)):
                if m[i] == solved_n:
                    m[i] = rules[solved_n]['regexp']
        solved = True
        for m in rule['matches']:
            if [str(_) for _ in m] != m:
                solved = False
        if solved:
            rules[n]['regexp'] = '('+'|'.join([''.join(m) for m in rule['matches']])+')'
            to_update.append(n)

from pprint import pprint
# pprint(rules)
regexp = '^' + rules[0]['regexp'] + '$'


import re
s = 0
for line in open('input.txt').readlines():
    if re.match(regexp, line):
        s+=1
print(s)