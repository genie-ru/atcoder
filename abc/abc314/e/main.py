N,M = map(int,input().split())
CPS = []

for i in range(N):
  li =tuple(map(int,input().split()))
  sl = [0] * (M+1)
  for j in li[2:]:
    sl[j] +=1
  CPS.append((li[0],li[1],sl))
CPS = sorted(CPS)

EM = []
for i in range(M):
  tm = 1e7
  for c,p,s in CPS:
    if c > tm:
      break
    e = (c*p + sum(EM[j] * s[i-j] for j in range(i))) / (p-s[0])
    # print(i,c,p,e)
    if e < tm:
      tm = e
  EM.append(tm)
print(tm)