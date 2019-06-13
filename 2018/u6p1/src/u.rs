use std::collections::*;

type Coord = (isize, isize);

fn to_coordinate(line: &str) -> Coord {
    scan_fmt!(line, "{d}, {d}", isize, isize).unwrap()
}

fn manhattan_distance(p: Coord, q: Coord) -> isize {
    (p.0 - q.0).abs() + (p.1 - q.1).abs()
}

fn closest(point: Coord, coords: &Vec<Coord>) -> Option<Coord> {
    let (winner, _) = coords.iter().fold((None, 1000), |(prev, d), test| {
        let pq = manhattan_distance(point, *test);
        if pq < d {
    //        println!("{} < {} shorter {:?} < {:?}", pq, d, prev, test);
            (Some(*test), pq)
        } else if pq == d {
    //        println!("{} = {} eq {:?} = {:?}", pq, d, prev, test);
            (None, pq)
        } else {
    //        println!("{} > {} gt {:?} > {:?}", pq, d, prev, test);
            (prev, d)
        }
    });
    //println!("winner: {:?} on point: {:?}", winner, point);
    winner
}

fn is_edged(p: Coord, e1: Coord, e2: Coord) -> bool {
    let b = p.0 == e1.0 || p.1 == e1.1 || p.0 == e2.0 || p.1 == e2.1;
    //println!("{} {:?} {:?} {:?}", b, p, e1, e2);
    b
}

pub fn solve(text: String) -> String {
    let coords: Vec<Coord> = text.lines().map(|c| to_coordinate(c)).collect();
    let le = coords.iter().min_by_key(|(x, _)| x.clone()).unwrap().0;// - 1;
    let ri = coords.iter().max_by_key(|(x, _)| x.clone()).unwrap().0;// + 1;
    let to = coords.iter().min_by_key(|(_, x)| x.clone()).unwrap().1;// - 1;
    let bo = coords.iter().max_by_key(|(_, x)| x.clone()).unwrap().1;// + 1;
    let tl: Coord = (le, to);
    let br: Coord = (ri, bo);

    let mut edged: HashSet<Coord> = HashSet::new();
    let mut size: HashMap<Coord, i32> = HashMap::new();
    //let mut winners: HashMap<Coord, Coord> = HashMap::new();

    for coords_x in le..(ri + 1) {
        for coords_y in to..(bo + 1) {
            let point = (coords_x, coords_y);
            //println!("p: {:?}", point);
            match closest(point, &coords) {
                Some(winner) => {
                    size.entry(winner).and_modify(|s| *s += 1).or_insert(1);
                    if is_edged(point, tl, br) {
                        edged.insert(winner);
                    }
                    //winners.entry(point).or_insert(winner);
                }
                None => {
                    /*
                    if is_edged(point, tl, br) {
                        //println!("edge: {:?} contested", point);
                    }
                    winners.entry(point).or_insert( (0,0) );
                    */
                }
            }
        }
    }

    let t = size.iter()
        .filter(|(k,_)| edged.insert(**k) )
        .max_by_key(|(_,v)| v.clone() );
    
    let wn = c_name(*t.unwrap().0);

    //let co: Vec<(usize,_)> = coords.iter().enumerate().collect();
    //println!("{:?}", co);
    /*
    for coords_x in le..(ri + 1) {
        for coords_y in to..(bo + 1) {
            let point = (coords_x, coords_y);
            let p = winners.get(&point).unwrap();
            let n = c_name(*p);
            print!("{}",n);
            if coords_y == bo {
                print!("\r\n");
            }
        }
    }
    */

    println!("{},{} {},{}", le, to, ri, bo);
    format!("{} {:?}", wn, t)
}

fn c_name(p: Coord) -> char {
    match p {
        (337, 150)  => 'a',
        (198, 248)  => 'b',
        (335, 161)  => 'c',
        (111, 138)  => 'd',
        (109, 48)   => 'e',
        (261, 155)  => 'f',
        (245, 130)  => 'g',
        (346, 43)   => 'h',
        (355, 59)   => 'i',
        (53, 309)   => 'j',
        (59, 189)   => 'k',
        (325, 197)  => 'l',
        (93, 84)    => 'm',
        (194, 315)  => 'n',
        (71, 241)   => 'o',
        (193, 81)   => 'p',
        (166, 187)  => 'q',
        (208, 95)   => 'r',
        (45, 147)   => 's',
        (318, 222)  => 't',
        (338, 354)  => 'u',
        (293, 242)  => 'v',
        (240, 105)  => 'x',
        (284, 62)   => 'y',
        (46, 103)   => 'z',
        (59, 259)   => 'A',
        (279, 205)  => 'B',
        (57, 102)   => 'C',
        (77, 72)    => 'D',
        (227, 194)  => 'E',
        (284, 279)  => 'F',
        (300, 45)   => 'G',
        (168, 42)   => 'H',
        (302, 99)   => 'I',
        (338, 148)  => 'J',
        (300, 316)  => 'K',
        (296, 229)  => 'L',
        (293, 359)  => 'M',
        (175, 208)  => 'N',
        (86, 147)   => 'O',
        (91, 261)   => 'P',
        (188, 155)  => 'Q',
        (257, 292)  => 'R',
        (268, 215)  => 'S',
        (257, 288)  => 'T',
        (165, 333)  => 'U',
        (131, 322)  => 'C',
        (264, 313)  => 'X',
        (236, 130)  => 'Y',
        (98, 60)    => 'Z',
        _           => '•'
    }
}
