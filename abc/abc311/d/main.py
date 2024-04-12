n,m = map(int,input().split())
S = [input() for _ in range(n)]
C = [[0]*m for _ in range(n)]
V = [(1,1)]
C[1][1] = 2
while V:
  x,y = V.pop()
  for p,q in ((-1,0),(0,-1),(1,0),(0,1)):
    i,j = x,y
    while S[i+p][j+q] == '.':
      C[i][j] = max(C[i][j],1)
      i,j = i+p,j+q
    if C[i][j] < 2:
      C[i][j] = 2
      V.append((i,j))
print(sum(C[i][j] > 0 for i in range(n) for j in range(m)))