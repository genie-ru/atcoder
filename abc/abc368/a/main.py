N, K = map(int, input().split())
A = list(map(int,input().split()))

B = []

for i in A[N-K:N+1]:
    B.append(i)
for i in A[0:N-K]:
    B.append(i)
print(*B)