pub fn solve(lines: String) -> String {

    fn to_u32(s: &str) -> i32 { s.parse().unwrap() }

    lines.lines() .map(to_u32)
        .fold(0, |acc, x| acc + x)
        .to_string()
}