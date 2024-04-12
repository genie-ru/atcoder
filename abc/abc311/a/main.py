n = int(input())
s = input()
a = 0
b = 0
c = 0
for i in range(n):
    if s[i] == "A":
        a += 1
    elif s[i] == "B":
        b += 1
    else:
        c += 1
    if a > 0 and b > 0 and c > 0:
        print(i + 1)
        break