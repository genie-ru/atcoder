N, *rest = map(int, input().split())
if len(rest) < N:
    rest += list(map(int, input().split()))
A = rest[:N]

operations = 0

while sum(a > 0 for a in A) > 1:
    A.sort(reverse=True)
    A[0] -= 1
    A[1] -= 1
    operations += 1

print(operations)
