N, P, Q = map(int, input().split())
D = list(map(int, input().split()))
D_min = min(D)
print(min(P, Q + D_min))
