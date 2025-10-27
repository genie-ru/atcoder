def status(i, M):
  return "OK" if i <= M else "Too Many Requests"

def main():
  N, M = map(int, input().split())
  print(*map(lambda i: status(i, M), range(1, N + 1)), sep="\n")

if __name__ == "__main__":
  main()