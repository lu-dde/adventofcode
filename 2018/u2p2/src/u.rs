
pub fn solve(text: String) -> String {
    let c1 = text.lines();

    fn find_distance<'a>(pair: (&'a str, &'a str)) -> (&'a str,&'a str,usize) {
       let (p,q) = pair; 
       let r = p.chars().zip(q.chars()).fold(0,|acc, (p,q)|acc + (p == q) as usize );
       (p,q,r)
    }

    let mut best: (&str,&str,usize) = (&"".to_string(),&"".to_string(),0);
    let mut i = 0;
    for e in c1 {
        i += 1;
        let c2 = text.lines().skip(i);
        for f in c2 {
            if e == f {
                continue;
            }
            let (p,q,r) = find_distance((e,f));
            if r > best.2 {
                best = (p,q,r);
            }
        }
    }

    let (p,q,r) = best;

    let s: Vec<char> = p.chars().zip(q.chars())
        .filter(|(a,b)| a == b)
        .map(|(a,_)|a)
        .collect();

    let mut key = String::new();
    for c in s { key.push(c); }

    format!("{} {} {} {}", key, r, p, q)
}
