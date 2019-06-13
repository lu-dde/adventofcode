use std::collections::HashSet;

fn eq_opposite_case(a: char, b: char) -> bool {
    a != b && a.eq_ignore_ascii_case(&b)
}

pub fn solve(text: String) -> String {
    let c1 = text.chars();

    let mut taken = HashSet::new();

    let min = c1
        .into_iter()
        .filter(|c| c.is_ascii_alphabetic())
        .map(|c| c.to_ascii_lowercase())
        .filter(|c| taken.insert(c.clone()))
        .map(|skip_c| {
            let length = text
                .chars()
                .filter(|c| !skip_c.eq_ignore_ascii_case(&c))
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
                .len();
            println!("{} {}", skip_c, length);
            length
        })
        .min()
        .unwrap();

    format!("{:?}", min)
}
