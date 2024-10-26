use std::collections::HashSet;
use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();

    let first_line = lines.next().unwrap().unwrap();
    let mut iter = first_line.split_whitespace();
    let n: i64 = iter.next().unwrap().parse().unwrap();
    let m: usize = iter.next().unwrap().parse().unwrap();
    
    let mut knights: Vec<(i64, i64)> = Vec::with_capacity(m);
    for _ in 0..m {
        let line = lines.next().unwrap().unwrap();
        let mut iter = line.split_whitespace();
        let x: i64 = iter.next().unwrap().parse().unwrap();
        let y: i64 = iter.next().unwrap().parse().unwrap();
        knights.push((x, y));
    }
    
    let mut attacked = HashSet::new();
    
    let mut occupied = HashSet::new();
    for &(x, y) in &knights {
        occupied.insert((x, y));
    }
    
    let moves = [
        (2, 1), (1, 2), (-1, 2), (-2, 1),
        (-2, -1), (-1, -2), (1, -2), (2, -1)
    ];
    
    for &(x, y) in &knights {
        for &(dx, dy) in &moves {
            let nx = x + dx;
            let ny = y + dy;
            if nx >= 1 && nx <= n && ny >= 1 && ny <= n {
                attacked.insert((nx, ny));
            }
        }
    }
    
    for pos in &occupied {
        attacked.remove(pos);
    }
    
    let total = n * n;
    let unsafe_count = (attacked.len() + occupied.len()) as i64;
    let safe_count = total - unsafe_count;
    
    println!("{}", safe_count);
}