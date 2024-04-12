use proconio::input;

fn main() {
    input! {
        a: usize,
        b: usize
    }
    if (b - a) != 1 {
        println!("No")
    } else if (a != 3) && (a != 6) {
        println!("Yes")
    } else {
        println!("No")
    }
}
