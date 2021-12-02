horizontal = 0
vertical = 0
aim = 0

with open('input.txt', 'r') as f:
    for line in f:
        ## split line by space and store in list
        line = line.split()
        if line[0] == "forward":
            horizontal += int(line[1])
            vertical += aim * int(line[1])
        elif line[0] == "down":
            aim += int(line[1])
        elif line[0] == "up":
            aim -= int(line[1])

    ##multiply horizontal and vertical to get total distance
    total = horizontal * vertical
    print(total)


        


