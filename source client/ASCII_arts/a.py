import os

for filename in os.listdir("."):
    if not filename.endswith(".py"):
        fle = ""
        with open(filename,'r') as f:
            for line in f.readlines():
                fle += " " + line
        with open(filename, 'w') as f:
            f.write(fle)