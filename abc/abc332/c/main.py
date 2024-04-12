N, M = map(int, input().split())
S = input()

logo = 0
M_status = M
logo_status = logo
for i in range(len(S)):
    if S[i] == "0":
        M_status = M
        logo_status = logo
    elif S[i] == "1":
        if M_status == 0 and logo_status == 0:
            logo += 1
        elif M_status > 0:
            M_status -= 1
        elif logo_status > 0:
            logo_status -=1
    elif S[i] == "2":
        if logo_status == 0:
            logo += 1
        else:
            logo_status -= 1
            
print(logo)