count = 0
previous = 0

with open('input.txt', 'r') as f:
    for line in f:
        if previous == 0:
            previous = int(line)
        else:
            if int(line) > previous:
                count += 1
            previous = int(line)

print(count)



