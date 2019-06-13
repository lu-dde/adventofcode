use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::LinkedList;


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

    let mut hm: HashMap<i32, LinkedList<i32>> = HashMap::new();
    let start = (HashSet::new(),HashSet::new());
    let (passed, invalid) = c1.map(pr).fold(start, |(mut passed,mut invalid), (id,p,q)| {
        passed.insert(id.clone());
        for i in to_sq(p,q) {
            hm.entry(i)
                .and_modify(|e| {
                    if e.len() > 1 {
                        invalid.insert(id.clone());
                    } else if e.len() == 1 {
                        invalid.insert(id.clone());
                        invalid.insert(e.front().unwrap().clone());
                    }
                    e.push_back(id.clone())
                } )
                .or_insert_with(|| {
                    let mut list: LinkedList<i32> = LinkedList::new();
                    list.push_back(id);
                    list
                } );
        }
        (passed,invalid)
    });

    let dif = passed.difference(&invalid);

    for x in dif {
        println!("{}", x);
    }

    format!("h")
}
