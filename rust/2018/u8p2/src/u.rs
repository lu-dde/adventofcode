//use rand::Rng;
use std::cmp::min;

#[derive(Debug)]
enum Action {
    ReadHead,
    Children( (i32,i32,i32) ),
    Eat(i32),
    EndNode,
}

pub fn solve(text: String) -> String {
    let mut children: Vec<_> = vec![];
    let mut meta: Vec<_> = vec![];
    let mut sum: Vec<i32> = vec![];
    let mut action: Action = Action::ReadHead;

    let mut t: std::iter::Peekable<_> = text
        .split_whitespace()
        .map(str::parse::<i32>)
        .filter_map(Result::ok)
        .peekable();

    'outer: loop {

        println!();
        println!();
        println!("children  {:?}", children);
        println!("meta      {:?}", meta);
        println!("sum       {:?}", sum);
        println!("action    {:?}", action);

        match action {
            Action::ReadHead => {
                if t.peek().is_some() {
                    let nr_of_children = t.next().unwrap();
                    let child = (0,min(1,nr_of_children),nr_of_children);
                    children.push( child );
                    let nr_of_meta = t.next().unwrap();
                    meta.push(nr_of_meta);

                    action = Action::Children( child );
                } else {
                    panic!("ok no next");
                }
            }
            Action::Children( (id,curr_child,child_count) ) if curr_child == child_count => {
                let meta = meta.pop().unwrap();
                action = Action::Eat(meta);
            }
            Action::Children( (id,curr_child,child_count) ) => {
                action = Action::ReadHead;
            }
            Action::Eat(0) => {
                action = Action::EndNode;
            }
            Action::Eat(num) => {
                let n = t.next().unwrap();
                sum.push(n);
                action = Action::Eat(num - 1);
            }
            Action::EndNode => {
                if let Some( (id,curr_child,child_count) ) = children.pop() {
                    if curr_child <= child_count {
                        children.push( (id,curr_child+1,child_count) );
                        action = Action::Children( (id,curr_child+1,child_count) );
                    } else {
                        action = Action::EndNode;
                    }
                } else {
                    break 'outer;
                }
            }
        }
    }

    sum.sort();

    println!();
    println!();
    println!("children  {:?}", children);
    println!("meta      {:?}", meta);
    //println!("sum       {:?}", sum);
    println!("sum       {:?}", sum.iter().fold(0, |x,y| x + y));
    println!("action    {:?}", action);

    format!("TREE SUM {:?}", 0)
}
