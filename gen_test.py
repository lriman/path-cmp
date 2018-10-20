import random
import math
total_pathes = 1000
total_points = 1000

min_speed = 10
max_speed = 100

max_dist = 1000
min_dist = 10

path_list = []

angles = (0, 45, 90, 135, 180, 225, 270, 315)

for path_id in range(total_pathes):
    x = random.randint(-1000, 1000)
    y = random.randint(-1000, 1000)
    t = random.randint(100, 1000)
    point = (path_id, x, y, t)
    for point_id in range(total_points):

        path_list.append(point)
        speed = random.randint(min_speed, max_speed)
        distance = random.randint(min_dist, max_dist)
        delta_t = distance / speed

        alpha = (math.pi * random.choice(angles)) / 180

        delta_x = distance*math.cos(alpha)
        delta_y = distance*math.sin(alpha)

        x = point[1] + delta_x
        y = point[2] + delta_y
        t = point[3] + delta_t
        point = (path_id, x, y, t)

s_list = sorted(path_list, key=lambda x: x[3])

f = open("test_2.list", "w")
for r in s_list:
    row = str(r[0]) + " " + str(r[1]) + " " + str(r[2]) + " " + str(r[3]) + "\n"
    f.write(row)

