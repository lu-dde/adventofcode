use std::collections::HashSet;

pub fn solve(lines: String) -> String {

    fn to_u32(s: &str) -> i32 { s.parse().unwrap() }

    let mut reached = HashSet::new();

    let mut i = lines.lines().map(to_u32).cycle();

    let mut acc = 0;
    loop {
        match i.next() {
            Some(n) => {
                acc += n;
                if ! reached.insert(acc) {
                    break
                }
            }
            _ => { break }
        }
    }

    acc.to_string()
}