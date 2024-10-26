n,c = map(int, input().split( ))
t =list(map(int, input().split( )))
pre =-1*c
candy = 0
for i in range(n):
    if t[i] >= pre + c:
        candy = candy + 1
        pre = t[i]
print(candy)