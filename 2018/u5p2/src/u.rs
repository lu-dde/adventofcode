fn eq_opposite_case(a: char, b: char) -> bool {
    a != b && a.eq_ignore_ascii_case(&b)
}

pub fn solve(text: String) -> String {
    let c1 = text.chars();

    let f: String = c1
        .filter(|c| c.is_ascii_alphabetic() )
        .rev()
        .fold(vec![], |mut f, c| {
            let option_fpop = f.pop();
            match option_fpop {
                Some(cC) => {
                    if !eq_opposite_case(cC, c) {
                        f.push(cC);
                        f.push(c);
                    }
                }
                None => {
                    f.push(c);
                }
            }
            f
        })
        .into_iter()
        .rev()
        .collect();

    format!("{}", f.len())
}
