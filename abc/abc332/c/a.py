N, M = map(int, input().split())
S = input()

logo = 0
M_status = M

for char in S:
    if char == "0":
        M_status = M  # 無地のTシャツをリセット
    elif char == "1":
        if M_status > 0:
            M_status -= 1
        else:
            logo += 1
    elif char == "2":
        if M_status > 0:
            M_status -= 1
        else:
            logo += 1

print(logo)
