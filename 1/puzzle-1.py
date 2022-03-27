import os

counter = 0
depth = None
last_depth = None

file = open(os.path.join(os.getcwd(), "input-puzzle-1"), "r")
lines = file.readlines()

for count, value in enumerate(lines):
    if count == 0:
        next
    depth = value
    last_depth = lines[count-1]
    if depth > last_depth:
        counter+=1

print(counter)
