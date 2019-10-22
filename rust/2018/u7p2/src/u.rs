use std::cmp;
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

fn get_cost(c: char) -> (char,i32) {
    (c, (c as i32) - 4)
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

    let mut total_time = 0;
    let mut workers: Vec<(char, i32)> = vec![];
    let max_workers = 5;

    let first = remaining
        .iter()
        .filter(|k| !disabled.contains_key(&k))
        .min()
        .unwrap();
    workers.push(get_cost(*first));

    let mut answer: String = "".to_string();

    let mut ready: Vec<char> = vec![];
    while !workers.is_empty() {
        workers.sort_by_key(|(c, i)| *i * 100 + (*c as i32));
        let (enabled, time_passed) = workers.remove(0);
        workers = workers
            .into_iter()
            .map(|(c, t)| (c, t - time_passed))
            .collect();
        total_time += time_passed;

        println!(
            "popped {} at time {} after {}s",
            enabled, total_time, time_passed
        );

        answer.push(enabled);

        let mut new_ready: Vec<char> = deps.iter()
            .filter(|(k, _)| *k == enabled)
            .filter_map(|(k, v)| {
                let set = disabled.get_mut(v).unwrap();
                set.remove(k);
                if set.is_empty() {
                    println!("     enabled: {}", v);
                    disabled.remove(v);
                    Some(v)
                } else {
                    None
                }
            }).cloned().collect();
        ready.append(&mut new_ready);

        ready.sort();

        println!("     READY: {:?}", ready);
        let min = cmp::min(ready.len(), max_workers - workers.len());
        for _ in 0..min {
            let new_char = ready.remove(0);
            workers.push(get_cost(new_char));
        }
        println!("     took: {} state: {:?}", min, workers);
    }

    format!("total_time {}", total_time)
}
