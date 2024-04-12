function Main(input) {
    const [[N], P, [Q], ...AB] = input.split("\n").map(e => e.split(" ").map(Number));
    const indexMap = new Map();
    P.forEach((element, index) => {
        indexMap.set(element, index);
    });

    const ans = [];
    for (let i = 0; i < Q; i++) {
        const [A, B] = AB[i];
        const a = indexMap.get(A);
        const b = indexMap.get(B);
        ans.push(a < b ? A : B);
    }
    console.log(ans.join("\n"));
}

Main(require("fs").readFileSync(0) + "");
