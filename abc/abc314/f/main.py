class UnionFind:
    def __init__(self, n):
        self.n = n
        self.parents = [-1] * n
    
    def root(self, x):
        pars = self.parents
        while pars[x] >= 0:
            x = pars[x]
        
        return x

    def unite(self, x, y):
        x = self.root(x)
        y = self.root(y)
        if x == y: return

        if self.parents[x] > self.parents[y]:
            x, y = y, x
        
        self.parents[x] += self.parents[y]
        self.parents[y] = x

        return x
    
    def size(self, x):
        return -self.parents[self.root(x)]
    
    def same(self, x, y):
        return self.root(x) == self.root(y)

import sys

MOD = 998244353

N = int(input())
matches = []
for _ in range(N - 1):
    p, q = map(int, sys.stdin.readline().split())
    p -= 1
    q -= 1
    matches.append((p, q))

max_ = N
inv = [1] * (max_ + 1)

for i in range(2, N + 1):
    inv[i] = -inv[MOD % i] * (MOD // i) % MOD

uf = UnionFind(N)
values = [0] * N
for p, q in matches:
    p_root = uf.root(p)
    q_root = uf.root(q)
    p_size = uf.size(p_root)
    q_size = uf.size(q_root)
    size_sum_inv = inv[p_size + q_size]
    p_value = p_size * size_sum_inv % MOD
    q_value = q_size * size_sum_inv % MOD
    values[p_root] += p_value
    values[q_root] += q_value
    parent = uf.unite(p_root, q_root)
    if parent == p_root:
        values[q_root] -= values[p_root]
    else:
        values[p_root] -= values[q_root]
    values[p_root] %= MOD
    values[q_root] %= MOD

children = [[] for _ in range(N)]
for i in range(N):
    if uf.parents[i] >= 0:
        children[uf.parents[i]].append(i)
    else:
        root = i

exps = values.copy()
for i in range(N):
    exps[i] %= MOD

stack = [root]
while stack:
    v = stack.pop()
    for child in children[v]:
        exps[child] += exps[v]
        exps[child] %= MOD
        stack.append(child)

print(*exps)
