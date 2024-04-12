N, M = map(int, input().split())
S = input()
C = list(map(int, input().split()))

color_positions = [[] for _ in range(M)]
for i in range(N):
    color_positions[C[i] - 1].append(i)

result = [''] * N
for i in range(M):
    shift_amount = 1 % len(color_positions[i])
    for j in range(len(color_positions[i])):
        original_index = color_positions[i][j]
        new_index = color_positions[i][(j + shift_amount) % len(color_positions[i])]
        result[new_index] = S[original_index]

print(''.join(result))
