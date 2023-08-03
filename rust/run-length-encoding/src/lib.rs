pub fn encode(source: &str) -> String {
    let mut s = source.chars().peekable();

    let mut res = String::new();
    let mut count = 0;
    while let Some(curr_char) = s.next() {
        count += 1;
        if s.peek() != Some(&curr_char) {
            if count > 1 {
                res.push_str(&count.to_string());
            }
            res.push(curr_char);
            count = 0;
        }
    }
    res
}

pub fn decode(source: &str) -> String {
    let mut res = String::new();
    let mut count = String::new();
    for c in source.chars() {
        if c.is_numeric() {
            count.push(c);
        } else {
            if !count.is_empty() {
                let repeat = count.parse::<usize>().unwrap();
                res.push_str(&c.to_string().repeat(repeat));
                count.clear();
            } else {
                res.push(c);
            }
        }
    }

    res
}
