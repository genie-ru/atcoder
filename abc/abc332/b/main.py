K, G, M = map(int, input().split())

x, y = 0, 0

for i in range(K):
    if x == G:
        x = 0
    elif y == 0:
        y = M

    else:
        z = min(y, G-x)
        x, y = x + z, y -z

print(str(x)+ " " + str(y))