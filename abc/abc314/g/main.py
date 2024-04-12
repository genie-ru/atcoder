import heapq
N, M, H = map(int, input().split())
D = [0]*(M+1) # current total damage
ans_ = [0]*(N+1)
amulet = 0 # need how many ? 0 at first
pql = [(0, i) for i in range(1, M+1)] # (-d, mid)
which = [0]*(M+1) # 0 is smaller 1 is amulet
heapq.heapify(pql)
curD = 0
pqr = []                              # (d, mid)
for i in range(1, N+1):
    A, B = map(int, input().split())
    pd = D[B]
    D[B] += A
    while pql and -pql[0][0] < D[pql[0][1]]:
        heapq.heappop(pql)
    if pql and D[B] <= -pql[0][0]:
        heapq.heappush(pql, (-D[B], B))
        which[B] = 0
        curD += A
        while curD > H - 1:
            while -pql[0][0] < D[pql[0][1]]:
                heapq.heappop(pql)
            d, mid = heapq.heappop(pql)
            heapq.heappush(pqr, (-d, mid))
            curD += d
            which[mid] = 1
            amulet += 1
    else:
        heapq.heappush(pqr, (D[B], B))
        if which[B] == 0:
            curD -= pd
            which[B] = 1
            amulet += 1
        while True:
            while pqr and pqr[0][0] < D[pqr[0][1]]:
                heapq.heappop(pqr)
            if pqr and curD + pqr[0][0] <= H - 1:
                d, mid = heapq.heappop(pqr)
                heapq.heappush(pql, (-d, mid))
                curD += d
                which[mid] = 0
                amulet -= 1
            else:
                break
    ans_[i] = amulet
ans = [0]*(M+1)
pm = ans_[N]
for j in range(pm, M+1):
    ans[j] = N
for i in range(1, N)[::-1]:
    if ans_[i] != pm:
        for j in range(ans_[i], pm):
            ans[j] = i
        pm = ans_[i]
print(*ans)
