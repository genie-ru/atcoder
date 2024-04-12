function Main(input) {
    input = input.split("\n").map(e => e.trim());
    const n = Number(input[0]);
    const s = input[1];
    let ok = [0, 0, 0];
    for (let i = 0; i < n; i++) {
        const str = s[i];
        if (str === "A") {
            ok[0]++;
        } else if (str === "B") {
            ok[1]++;
        } else {
            ok[2]++;
        }
        if(ok[0] * ok[1] * ok[2] !== 0) {
            console.log(i + 1);
            break;
        }
    }
}
Main(require("fs").readFileSync("/dev/stdin", "utf8"));