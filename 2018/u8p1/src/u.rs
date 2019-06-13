#![feature(integer_atomics)]
    
use std::sync::atomic::*;

static total_cost: AtomicI32 = AtomicI32::new(0);

#[derive(Debug)]
struct Node {
    meta: Vec<i32>,
    nr: i32,
    children: Box<Vec<Node>>,
    to_parse: Box<Vec<i32>>,
}

impl Node {

    pub fn new(to_parse: &[i32]) -> Node {
        let (meta,c,children) = Node::parse_split(to_parse);
        Node {
            meta: meta,
            nr: c,
            to_parse: Box::new(children),
            children: Box::new(vec![]),
        }
    }

    pub fn parse_split(to_parse: &[i32]) -> (Vec<i32>, i32, Vec<i32>) {
        let (node, rest) = to_parse.split_at(2);
        let c = node[0];
        let m = node[1];

        let children;
        let meta;
        if c == 0 {
            let (l, r) = rest.split_at(m as usize);
            children = r;
            meta = l;
        } else {
            let (l, r) = rest.split_at(rest.len() - (m as usize));
            children = l;
            meta = r;
        }
        (meta.to_vec(),c,children.to_vec())
    }

    pub fn expand(&mut self) {
        let t = &mut self.to_parse;
        for _ in 0..self.nr {
            let (m,n,c) = Node::parse_split(&t);
            t.drain(..);
            for a in (&c).iter() {
                t.push(*a);
            }
            for &cost in (&m).iter() {
                total_cost.fetch_add(cost,Ordering::Relaxed);
            }
            //println!("{} {}", n, total_cost.load(Ordering::Relaxed));
            let mut n = Node { meta: m, nr: n, children: Box::new(vec![]), to_parse: Box::new(c) };
            n.expand();
            self.children.push( n );
        }
        drop(self.to_parse);
    }

}

pub fn solve(text: String) -> String {
    let numbers: Vec<_> = text
        .split_whitespace()
        .map(|s| s.parse::<i32>() )
        .filter_map(Result::ok)
        .collect();

    // gör en vänster hänt consumer ät bara framåt i O(n)
    // skippa träd. det är inget träd!
    let mut n = Node::new(numbers.as_slice());
    n.expand();

    format!("TREE SUM {:?}", 0)
}