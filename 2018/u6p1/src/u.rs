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
            (Some(*test), pq)
        } else if pq == d {
            (None, pq)
        } else {
            (prev, d)
        }
    });
    //println!("    winner: {:?}", winner);
    winner
}

fn is_edged(p: Coord, e1: Coord, e2: Coord) -> bool {
    p.0 == e1.0 || p.1 == e1.1 || p.0 == e2.0 || p.1 == e2.1
}

pub fn solve(text: String) -> String {
    let coords: Vec<Coord> = text.lines().map(|c| to_coordinate(c)).collect();
    let le = coords.iter().min_by_key(|(x, _)| x.clone()).unwrap().0; // - 1;
    let ri = coords.iter().max_by_key(|(x, _)| x.clone()).unwrap().0; // + 1;
    let to = coords.iter().min_by_key(|(_, x)| x.clone()).unwrap().1; // - 1;
    let bo = coords.iter().max_by_key(|(_, x)| x.clone()).unwrap().1; // + 1;
    let tl: Coord = (le, to);
    let br: Coord = (ri, bo);

    let mut edged: HashSet<Coord> = HashSet::new();
    let mut size: HashMap<Coord, i32> = HashMap::new();

    for coords_x in le..(ri + 1) {
        for coords_y in to..(bo + 1) {
            let point = (coords_x, coords_y);
            //println!("p: {:?}", point);
            match closest(point, &coords) {
                Some(winner) => {
                    size.entry(winner).and_modify(|s| *s += 1).or_insert(1);
                    if is_edged(point, tl, br) {
                        edged.insert(winner);
                        let d = manhattan_distance(point, winner);
                        println!("edge: {:?} won by: {:?}   {}", point, winner, d);
                    }
                }
                None => {
                    if is_edged(point, tl, br) {
                        println!("edge: {:?} contested", point);
                    }
                }
            }
        }
    }

    let t = size.iter()
        //.inspect(|p| println!("{:?}", p) )
        .filter(|(k,_)| edged.contains(&k) )
        .map(|(_,v)| v )
        //.inspect(|p| println!("passed {:?}", p) )
        .max();

    println!("{},{} {},{}", le, to, ri, bo);
    format!("{:?}", t)
}
