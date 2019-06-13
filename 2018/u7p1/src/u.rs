use std::collections::*;

type Dep = (char,char);

fn parse_deps(line: &str) -> Dep {
    match scan_fmt!( line, "Step {} must be finished before step {} can begin.", char, char) {
        Ok(t) => t,
        _ => panic!("No such line"),
    }
}

pub fn solve(text: String) -> String {
    let deps: Vec<Dep> = text.lines().map(|c| parse_deps(c) ).collect();

    let mut remaining: HashSet<char> = HashSet::new();
    let mut disabled: HashSet<char> = HashSet::new();

    for d in &deps {
        remaining.insert(d.0);
        remaining.insert(d.1);
        disabled.insert(d.1);
    }

    let mut ds: Vec<char> = vec![];
    loop {
        let enabled: char;
        match remaining.difference(&disabled).min() {
            Some(c) => enabled = *c,
            None => break,
        }
        println!("{:?}", enabled);
        ds.push(enabled);
        remaining.remove(&enabled);
        deps.iter().filter(|(k,_)| *k == enabled ).for_each(|(k,v)| {
            println!("enabled {} by {}", v, k);
            disabled.remove(v);
        });
    }

    for x in ds {
        print!("{}", x);
    }
    println!();

    format!("CABDFE {:?}", 0)
}