N, X, Y, Z = map(int, input().split())

if (X < Y and X <= Z <= Y) or (Y < X and Y <= Z <= X):
    print("Yes")
else:
    print("No")