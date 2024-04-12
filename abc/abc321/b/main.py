N, X = map(int, input().split())
A = list(map(int, input().split()))
A.append(-1)

for last in range(0, 101):
    B = A.copy()
    B[N-1] = last
    B.sort()
    sum = 0
    for i in range(1, N-1):
        sum += B[i]
    if sum >= X:
        print(last)
        exit()
print('-1')