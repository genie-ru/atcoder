import sys

def main():
  numbers = list(map(int, sys.stdin.readline().split()))
  numbers.sort()

  if numbers[0] * numbers[1] == numbers[2]:
    print("Yes")
  else:
    print("No")

if __name__ == "__main__":
  main()