use std::time::Instant;

type Square = (i32, (i32, i32), (i32, i32));

fn pr(line: &str) -> Square {
    // #1281 @ 752,17: 29x18
    // #1282 @ 570,125: 19x10
    let mut t: Square = (0, (0, 0), (0, 0));
    match scan_fmt!(line, "#{d} @ {d},{d}: {d}x{d}", i32, i32, i32, i32, i32) {
        Ok((sq_id, x, y, dx, dy)) => t = (sq_id, (x, y), (x + dx - 1, y + dy - 1)),
        _ => {}
    }
    t
}

fn overlaps(p: Square, q: Square) -> bool {
    let (sq_id1, (a_l, a_t), (a_r, a_b)) = p;
    let (sq_id2, (b_l, b_t), (b_r, b_b)) = q;
    sq_id1 != sq_id2 && !(a_r < b_l || b_r < a_l || a_b < b_t || b_b < a_t)
    // sq_id1 != sq_id2 && a_r >= b_l && b_r >= a_l && a_b >= b_t && b_b >= a_t
}

pub fn solve(text: String) -> String {
    let now = Instant::now();

    let c1 = text.lines();

    let mut passed = vec![];
    let sq: Vec<Square> = c1.map(pr).collect();
    for head in (&sq).iter() {
        let id = &head.0;
        match sq.iter().find(|&other| overlaps(*head, *other)) {
            None => {
                passed.push(id);
            }
            _ => {}
        }
    }

    println!("{}", now.elapsed().as_millis());

    format!("{:?}", passed)
}
