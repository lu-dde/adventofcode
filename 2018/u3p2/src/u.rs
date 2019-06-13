use std::collections::HashMap;
use std::collections::HashSet;

fn pr(line: &str) -> Vec<(i32, i32)> {
    // #1281 @ 752,17: 29x18
    // #1282 @ 570,125: 19x10
    let mut a: Vec<_> = vec![];
    if let Ok((sq_id, x, y, dx, dy)) =
        scan_fmt!(line, "#{d} @ {d},{d}: {d}x{d}", i32, i32, i32, i32, i32)
    {
        for i in x..(x + dx) {
            for j in y..(y + dy) {
                a.push((sq_id, (i * 1000) + j));
            }
        }
    }
    a
}

pub fn solve(text: String) -> String {
    let c1 = text.lines();

    let start = (HashSet::new(), HashMap::new());
    let (invalid, _) = c1.flat_map(pr).fold(
        start,
        |(mut invalid, mut coords), (square_id, coordinate)| {
            if let Some(old) = coords.insert(coordinate, square_id) {
                invalid.insert(square_id);
                invalid.insert(old);
            }
            (invalid, coords)
        },
    );

    let mut t = 860016;

    for x in invalid {
        t -= x;
    }

    format!("{}", t)
}
