# 文字列 S を入力として受け取る
S = input().strip()

# a, e, i, o, u を取り除いた文字列を生成
filtered_string = ''.join(char for char in S if char not in 'aeiou')

# 結果を出力
print(filtered_string)