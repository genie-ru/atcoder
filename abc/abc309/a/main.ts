import * as fs from "fs";

const inputs = fs.readFileSync(process.stdin.fd, "utf8").trim();

const [a, b] = inputs.split(" ").map(value => parseInt(value));

const next = Math.abs(a - b);

if (next > 1) {
    console.log("No");
}

if (next === 1) {
    if (a % 3 === 0 && b % 3 === 1) {
        console.log("No");
    } else {
        console.log("Yes");
    }
}