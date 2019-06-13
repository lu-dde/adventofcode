
pub fn solve(text: String) -> String {
    let c1 = text.lines();
    let c2 = text.lines().skip(1);

    fn find_distance<'a>(pair: (&'a str, &'a str)) -> (&'a str,&'a str,usize) {
       let (p,q) = pair; 
       let r = p.chars().zip(q.chars()).fold(0,|acc, (p,q)|acc + (p == q) as usize );
       println!("{:?}",(p,q,r));
       (p,q,r)
    }

    let (p,q,i): (&str,&str,usize) = c1.zip(c2)
        .map(find_distance)
        .max_by(|a,b| a.2.cmp(&b.2))
        .unwrap();

    let s: Vec<char> = p.chars().zip(q.chars())
        .filter(|(a,b)| a == b)
        .map(|(a,_)|a)
        .collect();

    let mut key = String::new();
    for c in s {
        key.push(c);
    }

    println!("{:?} {}", key, key.len());

    format!("{} = {} {} {}", p, q, i, p.len())
}
