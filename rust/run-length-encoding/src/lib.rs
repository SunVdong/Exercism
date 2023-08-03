pub fn encode(source: &str) -> String {
    let mut res = String::new();
    let mut count = 1;
    let mut chars = source.chars().peekable();

    while let Some(current_char) = chars.next() {
        if let Some(next_char) = chars.peek() {
            if current_char == *next_char {
                count += 1;
            } else {
                if count > 1 {
                    res.push_str(&count.to_string());
                }
                res.push(current_char);
                count = 1;
            }
        } else {
            if count > 1 {
                res.push_str(&count.to_string());
            }
            res.push(current_char);
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
