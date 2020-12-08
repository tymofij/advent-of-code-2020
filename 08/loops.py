code = [[s.split()[0], int(s.split()[1])] for s in  open('input.txt').readlines()]

def run(code):
    acc = 0
    visited = set()
    i = 0
    while i < len(code):
        if i in visited:
            return acc, True
        visited.add(i)
        instr, n = code[i]
        if instr == 'acc':
            acc += n
            i += 1
        elif instr == 'nop':
            i += 1
        elif instr == 'jmp':
            i += n
    return acc, False

acc, looped = run(code)
print('Accumulator at first loop:', acc)

for i in range(len(code)):
    instr, n = code[i]
    if instr == 'acc':
        continue
    code[i][0] = 'nop' if instr =='jmp' else 'jmp'
    acc, looped = run(code)
    if not looped:
        print('Final accumulator in fixed code:', acc)
        break
    code[i][0] = instr
