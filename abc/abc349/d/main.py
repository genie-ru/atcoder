def f(l, r, L, R):
    if L <= l and r <= R:
        return [(l, r)]
    m = (l + r) // 2
    if R <= m:
        return f(l, m, L, R)
    if m <= L:
        return f(m, r, L, R)
    return f(l, m, L, R) + f(m, r, L, R)

L, R = map(int, input().split())
ans = f(0, 1 << 60, L, R)
print(len(ans))
for l, r in ans:
    print(l, r)