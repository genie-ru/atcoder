function main(input) {
    const chars = input.trim().split("");
    let charCount = {};

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
require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
}).on('line', (line) => {
    main(line);
    process.exit();
});
