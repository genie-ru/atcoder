use std::io;

fn main() {
  let mut input = String::new();
  io::stdin().read_line(&mut input).expect("入力エラー");
  let x: f64 = input.trim().parse().expect("数値として解析できません");

  if x >= 38.0 {
    println!("1");
  } else if x >= 37.5 {
    println!("2");
  } else {
    println!("3");
  }
}