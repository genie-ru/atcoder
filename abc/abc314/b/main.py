*a,x=[{*map(int,t.split())}for t in open(0)]
L,m=[],99
for i,s in enumerate(a[2::2],1):
 if x&s:
  if len(s)==m:L+=[i]
  if len(s)<m:L,m=[i],len(s)
print(len(L))
print(*L)
