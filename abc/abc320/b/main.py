def max_palindrome_length(S):
    n = len(S)
    max_length = 0

    # dp[i][j]はS[i:j+1]が回文であるかを示す
    dp = [[False] * n for _ in range(n)]

    # すべての文字は単体で回文になる
    for i in range(n):
        dp[i][i] = True
        max_length = max(max_length, 1)

    # 2文字からなる回文をチェック
    for i in range(n - 1):
        if S[i] == S[i + 1]:
            dp[i][i + 1] = True
            max_length = max(max_length, 2)

    # 3文字以上からなる回文をチェック
    for length in range(3, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if S[i] == S[j] and dp[i + 1][j - 1]:
                dp[i][j] = True
                max_length = max(max_length, length)

    return max_length

# 入力文字列
S = input()  # 例: "BANANA"

# 最大回文長を求める
result = max_palindrome_length(S)
print(result)
