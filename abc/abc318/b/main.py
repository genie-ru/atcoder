N = int(input())
rects = []
for _ in range(N):
    A, B, C, D = map(int, input().split())
    rects.append((A, B, C, D))

grid = [[0 for _ in range(101)] for _ in range(101)]
for rect in rects:
    A, B, C, D = rect
    for i in range(A, B):
        for j in range(C, D):
            grid[i][j] = 1

S = 0
for row in grid:
    S += sum(row)
print(S)
