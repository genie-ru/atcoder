use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let mut numbers: Vec<i32> = input
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();
    
    numbers.sort(); // 数値をソートすることで、組み合わせを効率的に確認できる

    if numbers[0] * numbers[1] == numbers[2] {
        println!("Yes");
    } else {
        println!("No");
    }
}
