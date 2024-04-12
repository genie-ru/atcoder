function main(input: string): void {
    const chars: string[] = input.trim().split("");
    let charCount: { [key: string]: number } = {};

    // 各文字の出現回数をカウント
    for (const char of chars) {
        if (charCount[char]) {
            charCount[char]++;
        } else {
            charCount[char] = 1;
        }
    }

    // 唯一異なる文字を探し、その位置を出力
    for (let i = 0; i < chars.length; i++) {
        if (charCount[chars[i]] === 1) {
            console.log(i + 1); // 1-indexed にするために 1 を加える
            return;
        }
    }
}

// 標準入力からの読み込みを想定したコード
import * as readline from 'readline';

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.on('line', (line) => {
    main(line);
    rl.close();
});
