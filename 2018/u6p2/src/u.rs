use std::collections::*;

type Coord = (isize, isize);

fn to_coordinate(line: &str) -> Coord {
    scan_fmt!(line, "{d}, {d}", isize, isize).unwrap()
}

fn manhattan_distance(p: Coord, q: Coord) -> isize {
    (p.0 - q.0).abs() + (p.1 - q.1).abs()
}

fn in_region(point: Coord, coords: &[Coord], distance: isize) -> bool {
    coords
        .iter()
        .fold(0, |d, test| d + manhattan_distance(point, *test))
        < distance
}

pub fn solve(text: String) -> String {
    let coords: Vec<Coord> = text.lines().map(|c| to_coordinate(c)).collect();
    let le = coords.iter().min_by_key(|(x, _)| *x).unwrap().0; // - 1;
    let ri = coords.iter().max_by_key(|(x, _)| *x).unwrap().0; // + 1;
    let to = coords.iter().min_by_key(|(_, x)| *x).unwrap().1; // - 1;
    let bo = coords.iter().max_by_key(|(_, x)| *x).unwrap().1; // + 1;

    let mut region: HashSet<Coord> = HashSet::new();

    for coords_x in le..=ri {
        for coords_y in to..=bo {
            let point = (coords_x, coords_y);
            if in_region(point, &coords, 10000) {
                region.insert(point);
            }
        }
    }

    let print_matrix = true;
    if print_matrix {
        for coords_x in le..=ri {
            for coords_y in to..=bo {
                if region.contains(&(coords_x, coords_y)) {
                    print!("•")
                } else {
                    print!(" ")
                }
                if coords_y == bo {
                    println!();
                }
            }
        }
    }

    println!("{},{} {},{}", le, to, ri, bo);
    format!("{:?}", region.len())
}

fn c_name(p: Coord) -> char {
    match p {
        (337, 150) => 'a',
        (198, 248) => 'b',
        (335, 161) => 'c',
        (111, 138) => 'd',
        (109, 48) => 'e',
        (261, 155) => 'f',
        (245, 130) => 'g',
        (346, 43) => 'h',
        (355, 59) => 'i',
        (53, 309) => 'j',
        (59, 189) => 'k',
        (325, 197) => 'l',
        (93, 84) => 'm',
        (194, 315) => 'n',
        (71, 241) => 'o',
        (193, 81) => 'p',
        (166, 187) => 'q',
        (208, 95) => 'r',
        (45, 147) => 's',
        (318, 222) => 't',
        (338, 354) => 'u',
        (293, 242) => 'v',
        (240, 105) => 'x',
        (284, 62) => 'y',
        (46, 103) => 'z',
        (59, 259) => 'A',
        (279, 205) => 'B',
        (57, 102) => 'C',
        (77, 72) => 'D',
        (227, 194) => 'E',
        (284, 279) => 'F',
        (300, 45) => 'G',
        (168, 42) => 'H',
        (302, 99) => 'I',
        (338, 148) => 'J',
        (300, 316) => 'K',
        (296, 229) => 'L',
        (293, 359) => 'M',
        (175, 208) => 'N',
        (86, 147) => 'O',
        (91, 261) => 'P',
        (188, 155) => 'Q',
        (257, 292) => 'R',
        (268, 215) => 'S',
        (257, 288) => 'T',
        (165, 333) => 'U',
        (131, 322) => 'C',
        (264, 313) => 'X',
        (236, 130) => 'Y',
        (98, 60) => 'Z',

        (1, 1) => 'A',
        (1, 6) => 'B',
        (8, 3) => 'C',
        (3, 4) => 'D',
        (5, 5) => 'E',
        (8, 9) => 'F',

        /*
                (1, 1) => 'A',
                (6, 1) => 'B',
                (3, 8) => 'C',
                (4, 3) => 'D',
                (5, 5) => 'E',
                (9, 8) => 'F',
        */
        _ => '•',
    }
}
