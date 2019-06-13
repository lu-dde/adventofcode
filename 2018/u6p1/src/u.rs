fn to_coordinate(line: &str) -> (usize,usize) {
    match scan_fmt!(line, "{d}, {d}", usize, usize) {
        Ok(t) => t,
        _ => {panic!("failed to parse coordinate") }
    }
}

pub fn solve(text: String) -> String {
    let c1 = text.lines().map(|c| to_coordinate(c));


    format!("{:?}", c1)
}
