use std::io::{self, BufRead};
use std::collections::HashMap;

struct TrieNode {
    children: HashMap<u8, TrieNode>,
    min_value: i64,
}

impl TrieNode {
    fn new() -> Self {
        TrieNode {
            children: HashMap::new(),
            min_value: i64::MAX,
        }
    }
}

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    let n: usize = lines.next().unwrap().unwrap().trim().parse().unwrap();
    let mut s_list = Vec::with_capacity(n);
    for _ in 0..n {
        let s = lines.next().unwrap().unwrap();
        s_list.push(s.into_bytes());
    }

    let mut root = TrieNode::new();

    for k in 0..n {
        let s = &s_list[k];
        let len_k = s.len() as i64;
        let mut min_total_cost = len_k;

        let mut current_node = &root;
        let mut depth = 0i64;

        for &c in s {
            if let Some(next_node) = current_node.children.get(&c) {
                depth += 1;
                let candidate_cost = len_k + next_node.min_value;
                if candidate_cost < min_total_cost {
                    min_total_cost = candidate_cost;
                }
                current_node = next_node;
            } else {
                break;
            }
        }

        println!("{}", min_total_cost);

        let mut current_node = &mut root;
        depth = 0;
        for &c in s {
            current_node = current_node.children.entry(c).or_insert_with(TrieNode::new);
            depth += 1;
            let computed_value = len_k - 2 * depth;
            if computed_value < current_node.min_value {
                current_node.min_value = computed_value;
            }
        }
    }
}
