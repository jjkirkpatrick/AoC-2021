count = 0
readings = []

with open('input.txt', 'r') as f:
    for line in f:
        readings.append(int(line))

listlen = len(readings)
previous = 0

for i in range(listlen):
    window = readings[i:i+3]
    if len(window) == 3:
        if previous == 0:
            previous = sum(window)
        elif sum(window) > previous:
            count += 1
        previous = sum(window)

            
            

print(count)