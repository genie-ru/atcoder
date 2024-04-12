const main = (input: string) => {
    const n = input.split('').map(s => parseInt(s))
    for (let i = 0; i < n.length - 1; i++) {
      if (n[i] <= n[i + 1]) {
        console.log('No')
        return
      }
    }
    console.log('Yes')
  };
  main(require("fs").readFileSync("/dev/stdin", "utf8").trim());
  