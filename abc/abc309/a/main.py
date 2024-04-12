A, B = map(int, input().split())
if (A-B == 1) or (B-A == 1):
    if(A !=3) and (A != 6):
        print('Yes')
    else:
        print("No")
else:
    print("No")