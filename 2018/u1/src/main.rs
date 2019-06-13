use std::env;
use std::fs::File;
use std::io::prelude::*;

mod u;

fn main() {
    let mut res = None;
    let args: Vec<String> = env::args().collect();
    match args.len() {
        1 => {
            println!("atleast 1 argument");
        }
        2 => {
            let file_name = &args[1];
            println!("reading file:'{}'", file_name);
            let file = File::open(file_name);
            let mut buf = String::new();
            match file {
                Ok(mut file) => {
                    match file.read_to_string(&mut buf) {
                        Ok(_) => {
                            println!("buf: {:#?}", buf);
                            res = Some(u::solve(buf));
                        }
                        Err(why) => {
                            println!("ERROR: {:#?}", why);
                        }
                    }
                }
                Err(why) => {
                    println!("ERROR: {:#?}", why);
                }
            }
        }
        _ => {
            println!("unmatching arguments");
        }
    }

    match res {
        Some(answer) => {
            println!("anwser is: {:#?}", answer);
        }
        _ => {
            println!("unresolved");
        }
    }
}
