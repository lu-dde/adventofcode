use std::fs::File;
use std::io::prelude::*;

mod u;

fn main() {
    let file = File::open("input");
    let res = match file {
        Ok(mut f) => {
            let mut buf = String::new();
            match f.read_to_string(&mut buf) {
                Ok(_) => {
                    Some(u::solve(buf))
                }
                Err(why) => {
                    println!("ERROR: {:#?}", why);
                    None
                }
            }
        }
        Err(why) => {
            println!("ERROR: {:#?}", why);
            None
        }
    };

    match res {
        Some(answer) => {
            println!("anwser is: {:#?}", answer);
        }
        None => {
            println!("unresolved");
        }
    }
}
