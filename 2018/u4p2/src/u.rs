use chrono::prelude::*;
use std::collections::*;

fn parse_guard_number(s: &str) -> i64 {
    // #123 => 123
    let t = s.chars().next().map(|c| &s[c.len_utf8()..]);
    match t.unwrap_or("").parse::<i64>() {
        Ok(n) => n,
        _ => panic!("failed to parse!")
    }
}

//[1518-08-08 00:55] wakes up
//[1518-04-14 00:48] falls asleep
//[1518-11-20 00:03] Guard #2333 begins shift

#[derive(Debug, Hash, Eq, PartialEq, Ord, PartialOrd)]
enum GE {
    Shift(i64),
    Asleep,
    Wake,
}

fn guardevent_times(line: &str) -> (DateTime<Utc>, GE) {
    let t: GE;
    let i: DateTime<Utc>;
    match scan_fmt!(
        line,
        "[{}-{}-{} {}:{}] {} {}",
        i32,
        u32,
        u32,
        u32,
        u32,
        String,
        String
    ) {
        Ok((y, mo, d, h, mi, action, nr)) => {
            i = Utc.ymd(y, mo, d).and_hms(h, mi, 0);

            match action.as_ref() {
                "wakes" => t = GE::Wake,
                "falls" => t = GE::Asleep,
                "Guard" => t = GE::Shift(parse_guard_number(&nr)),
                _ => panic!("No such action"),
            }
        }
        _ => panic!("No such line"),
    }
    (i, t)
}

pub fn solve(text: String) -> String {
    let c1 = text.lines();

    let mut p: Vec<(DateTime<Utc>, GE)> = c1.map(guardevent_times).collect();
    p.sort_by(|(a, _), (b, _)| a.partial_cmp(b).unwrap());

    let q = p.iter().zip(p.iter().cycle().skip(1));

    let mut minute_count: HashMap<_,i64> = HashMap::new();

    // possible transitions
    //   Asleep    =>  Wake
    //   Shift(N)  =>  Asleep
    //   Shift(N)  =>  Shift(N)
    //   Wake      =>  Asleep
    //   Wake      =>  Shift(N)
    let mut current_guard_id: i64 = 0;
    for ((sleep_start, a), (sleep_to, _)) in q {
        match a {
            GE::Shift(n) => current_guard_id = *n,
            GE::Asleep => {
                let nr_of_minutes =
                    (*sleep_to - *sleep_start).num_minutes();
                let s = sleep_start.time();
                for offset in 0..nr_of_minutes {
                    minute_count
                        .entry((current_guard_id, s + chrono::Duration::minutes(offset) ))
                        .and_modify(|e| *e += 1)
                        .or_insert(1);
                }
            }
            _ => {}
        }
    }

    let ((gid,minute), count) = minute_count
        .iter()
        .max_by(|((_, _), x), ((_, _), y)| x.cmp(y)).unwrap();


    println!("guard {} minute {} count {}", gid, minute, count);
    format!("{:#?}", gid * minute.minute() as i64)
}
