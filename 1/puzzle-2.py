import os

counter = 0
last_total = None
total = None

file = open(os.path.join(os.getcwd(), "input-puzzle-1"), "r")
# clean the file
lines = file.readlines()
for count, value in enumerate(lines):
    lines[count] = value.strip()
    if (not value) or (value == "") or value == '\n':
        del lines[count]
        continue
    lines[count] = int(lines[count])

sums_lst = []


def get_sum(index):
    depth = lines[index]
    last_depth = lines[index-1]
    next_depth = lines[index+1]
    s = depth + last_depth + next_depth
    return s


def build_sums_lst():
    for count, value in enumerate(lines):
        if count == 0:
            continue
        if count + 1 == len(lines):
            break
        l = get_sum(count)
        #breakpoint()
        sums_lst.append(l)


def count_all():
    global counter
    build_sums_lst()
    for count, value in enumerate(sums_lst):
        if count == 0:
            continue
        total = value
        last_total = sums_lst[count-1]
        if total > last_total:
            counter+=1
            print(str(total), " (increased)")
        else:
            print(str(total))
        
count_all()
print(counter)
