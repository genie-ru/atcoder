N, S, K = map(int, input().split())
res = 0

for i in range(N):
    p, q = map(int, input().split())
    res = res + p * q

if res >= S:
    print(res)
else:
    print(res + K)
