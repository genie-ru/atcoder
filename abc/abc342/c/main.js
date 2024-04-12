let lines = require("fs").readFileSync("/dev/stdin", "utf8").trim().split("\n");
let N = Number(lines[0]);
let data = "abcdefghijklmnopqrstuvwxyz".split('');
let S = lines[1].split("");
let Q = Number(lines[2]);

let transformMap = new Map(data.map((c, i) => [c, c]));

for (let i = 0; i < Q; i++) {
  let [c, d] = lines[i + 3].split(" ");
  data.forEach((char, index) => {
    if (transformMap.get(char) === c) {
      transformMap.set(char, d);
    }
  });
}

for (let i = 0; i < N; i++) {
  S[i] = transformMap.get(S[i]);
}

console.log(S.join(""));
