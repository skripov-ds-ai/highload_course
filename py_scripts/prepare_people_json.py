import os
import csv
import json
import random

arr = []

path = os.path.dirname(__file__).replace("py_scripts", "db/migrations/data/people.csv")
with open(path, "r") as f:
    r = csv.reader(f, "excel", delimiter=",")
    for row in r:
        names = row[0].split()
        d = {"first_name": names[1], "last_name": names[0]}
        arr.append(d)
# print(arr)
random.shuffle(arr)

save_path = os.path.dirname(__file__).replace("py_scripts", "k6/people_names.json")
with open(save_path, "w") as f:
    json.dump(arr, f, indent=6)
