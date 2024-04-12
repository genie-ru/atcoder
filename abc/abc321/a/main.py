N = input()

prev_digit = float('inf')

for digit in N:
    digit = int(digit)
    if digit >= prev_digit:
        print('No')
        exit()
    prev_digit = digit

print('Yes')