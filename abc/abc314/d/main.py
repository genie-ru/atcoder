import sys

input = sys.stdin.readline
N = int(input())
S = list(input())
last_modified = [-1] * N
Q = int(input())
l = -1
u = -1
for i in range(Q):
    t, x, c = input().split()
    if t == "1":
        x = int(x) - 1
        S[x] = c
        last_modified[x] = i
    elif t == "2":
        l = i
    else:
        u = i

for i in range(N):
    if l < u and last_modified[i] < u:
        S[i] = S[i].upper()
    elif u < l and last_modified[i] < l:
        S[i] = S[i].lower()

print("".join(S))
