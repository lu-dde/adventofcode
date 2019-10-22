use std::collections::HashMap;

use regex::Regex;

fn pv(x: i32,y: i32) -> i32 {
    x * 1000 + y
}

fn pr(t: &str) -> (i32, (i32,i32), (i32,i32)) {
    let re = Regex::new(r"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)").unwrap();
    let c = re.captures(t).unwrap();
    let i1 = c.get(1).unwrap().as_str().parse::<i32>().unwrap();
    let i2 = c.get(2).unwrap().as_str().parse::<i32>().unwrap();
    let i3 = c.get(3).unwrap().as_str().parse::<i32>().unwrap();
    let i4 = c.get(4).unwrap().as_str().parse::<i32>().unwrap();
    let i5 = c.get(5).unwrap().as_str().parse::<i32>().unwrap();
    (i1,(i2,i3),(i2+i4,i3+i5))
}

fn to_sq(p1: (i32,i32),p2: (i32,i32)) -> Vec<i32> {
    let (x1,y1) = p1;
    let (x2,y2) = p2;
    let mut a: Vec<i32> = vec![];
    for i in x1..x2 {
        for j in y1..y2 {
            a.push(pv(i,j));
        }
    }
    a
}

pub fn solve(text: String) -> String {
    let c1 = text.lines();

    let mut hm = HashMap::new();

    let p: Vec<i32> = c1.map(pr).map(|(_,p,q)| to_sq(p,q) ).flatten().collect();
    for i in p {
      hm.entry(i).and_modify(|e| { *e += 1 } ).or_insert(1);
    }
    let k = hm.values().fold(0,|acc,i| acc + (*i > 1) as i32);

    let pa = hm.len();

    format!("{} {}", pa, k)
}
