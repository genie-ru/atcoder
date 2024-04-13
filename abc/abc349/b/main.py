from collections import Counter

S = input()

char_counts = Counter(S)


count_counts = Counter(char_counts.values())

for count in count_counts.values():
    if count != 0 and count != 2:
        print("No")
        break
else:
    print("Yes")