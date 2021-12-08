from __future__ import print_function
import sys
import os

def main():

    if len(sys.argv) != 3:
        return

    p = sys.argv[2]
    if not os.path.exists(p):
        return

    branch = sys.argv[1].split("/")
    if len(branch) == 0:
        print("WARNING: not appending issue number: unexpected branch name")
        return

    branch = branch[-1].split("-")
    if len(branch) == 0:
        print("WARNING: not appending issue number: unexpected branh name")
        return

    # should be 4 numbers
    if not branch[0].isdigit():
        print("WARNING: not appending issue number: unexpected branh name")
        return

    to_append = " (#"+branch[0]+')'
    #print(to_append)

    lines = []
    with open(p, "r") as f:
        for line in f.readlines():
            lines.append(line.strip())

    if len(lines) != 1:
        print("WARNING: not appending issue number: unexpected lines")
        return

    lines[-1] = lines[-1] + to_append
    with open(p,"w") as f:
        for line in lines:
            f.write(line+"\n")


if __name__ == "__main__":
    main()

