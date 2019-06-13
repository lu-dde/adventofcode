use std::collections::*;

type Dep = (char, char);

fn parse_deps(line: &str) -> Dep {
    match scan_fmt!(
        line,
        "Step {} must be finished before step {} can begin.",
        char,
        char
    ) {
        Ok(t) => t,
        _ => panic!("No such line"),
    }
}

pub fn solve(text: String) -> String {
    let deps: Vec<Dep> = text.lines().map(|c| parse_deps(c)).collect();

    let mut remaining: HashSet<char> = HashSet::new();
    let mut disabled: HashMap<char, HashSet<char>> = HashMap::new();

    for d in &deps {
        remaining.insert(d.0);
        remaining.insert(d.1);
        disabled.entry(d.1).or_default().insert(d.0);
    }

    let max_workers = 5;
    let mut workers: Vec<(char,i32)> = vec![];

    let mut ds: Vec<char> = vec![];
    loop {
        let enabled: char;
        match remaining
            .iter()
            .filter(|k| !disabled.contains_key(&k))
            .min()
        {
            Some(c) => enabled = *c,
            None => break,
        }
        print!("{}", enabled);
        ds.push(enabled);
        remaining.remove(&enabled);
        deps.iter()
            .filter(|(k, _)| *k == enabled)
            .for_each(|(k, v)| {
                let set = disabled.get_mut(v).unwrap();
                set.remove(k);
                if set.is_empty() {
                    disabled.remove(v);
                    //println!("enabled {} by {}", v, k);
                }
            });
    }

    println!();
    format!("CABDFE {:?}", 0)
}
