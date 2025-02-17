import sys

def missing_numbers(n, m, a):
    all_numbers = set(range(1, n + 1))
    a_set = set(a)
    missing = sorted(all_numbers - a_set)
    
    print(len(missing))
    if missing:
        print(" ".join(map(str, missing)))
    else:
        print("")  # 出力が0のときは空行を明示的に出力

if __name__ == "__main__":
    n, m = map(int, sys.stdin.readline().split())
    a = list(map(int, sys.stdin.readline().split()))
    missing_numbers(n, m, a)
