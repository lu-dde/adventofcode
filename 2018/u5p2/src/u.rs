use std::collections::HashSet;

fn eq_opposite_case(a: char, b: char) -> bool {
    a != b && a.eq_ignore_ascii_case(&b)
}

fn react(skip_char: char, text: &Vec<char>) -> usize {
    let length = text
        .iter()
        .filter(|c| !skip_char.eq_ignore_ascii_case(&c))
        .fold(vec![], |mut f, c| {
            let option_fpop = f.pop();
            match option_fpop {
                Some(c_c) => {
                    if !eq_opposite_case(c_c, *c) {
                        f.push(c_c);
                        f.push(*c);
                    }
                }
                None => {
                    f.push(*c);
                }
            }
            f
        })
        .len();
    length
}

pub fn solve(text: String) -> String {
    let c1: Vec<char> = text
        .chars()
        .filter(|c| c.is_ascii_alphabetic())
        .collect();

    let mut taken = HashSet::new();

    let min = (&c1)
        .iter()
        .map(|c| c.to_ascii_lowercase())
        .filter(|c| taken.insert(c.clone()))
        .map(|skip_char: char| react(skip_char, &c1))
        .min()
        .unwrap();

    format!("{}", min)
}
