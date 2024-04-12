# 横並び
# fa
H, W = map(int, input().split())
# 縦並び
A = [list(map(int, input().split())) for _ in range(H)]
B = [list(map(int, input().split())) for _ in range(H)]

from itertools import permutations


def move_cnt(tup):
    n = len(tup)
    result = 0
    for i in range(n - 1):
        for j in range(i + 1, n):
            if tup[i] > tup[j]:
                result += 1
    return result

ans = 120 * 120 + 1

# 高さのindexの順列のループ
for i_perm in permutations(range(H), H):
    # 幅のindexの順列のループ
    for j_perm in permutations(range(W), W):
        C = [[A[i][j] for j in j_perm] for i in i_perm]
        if C == B:
            ans = min(ans, move_cnt(i_perm) + move_cnt(j_perm))

if ans == 120 * 120 + 1:
    ans = -1
print(ans)
