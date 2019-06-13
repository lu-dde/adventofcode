#[macro_use]
extern crate scan_fmt;

use std::time::Instant;

use std::fs::File;
use std::io::prelude::*;

mod u;

fn main() {
    match File::open("input") {
        Ok(mut f) => {
            let mut buf = String::new();
            match f.read_to_string(&mut buf) {
                Ok(_) => {
                    let now = Instant::now();
                    println!("anwser is: {:#?}", u::solve(buf));
                    println!("{}ms", now.elapsed().as_millis());
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
