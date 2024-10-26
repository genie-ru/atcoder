use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let reader = stdin.lock();
    let mut row_mask: u8 = 0;
    let mut col_mask: u8 = 0;
    
    for (i, line) in reader.lines().take(8).enumerate() {
        let row = line.unwrap();
        
        for (j, ch) in row.chars().take(8).enumerate() {
            if ch == '#' {
                row_mask |= 1 << i;
                col_mask |= 1 << j;
            }
        }
    }

    let safe_rows = !row_mask;
    let safe_cols = !col_mask;
    let mut count = 0;
    
    for i in 0..8 {
        if safe_rows & (1 << i) != 0 {
            count += (safe_cols as u32).count_ones() as i32;
        }
    }
    
    println!("{}", count);
}