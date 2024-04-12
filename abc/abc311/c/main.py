import sys
input=sys.stdin.readline
n=int(input().strip())
a=list(map(int,input().strip().split(" ")))
o=[]
i=0
while True:
    o.append(i+1)
    j=i
    i=a[i]-1
    a[j]=0
    
    if a[i]==0:
        break
ou=o[o.index(i+1):len(o)]


print(len(ou))
print(" ".join(list(map(str,ou))))
    
        
