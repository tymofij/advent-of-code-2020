import re

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

# part 2 looping modification
for i in range(2,10):
    # rule 8 matches 42, 42 42, 42 42 42 and on
    rules[8]['matches'].append([42]*i)
    # rule 11 matches 42 31, 42 42 31 31, 42 42 42 31 31 31 and on
    rules[11]['matches'].append([42]*i + [31]*i)

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

regexp = '^' + rules[0]['regexp'] + '$'
print(sum([1 for line in open('input.txt').readlines() if re.match(regexp, line)]))
