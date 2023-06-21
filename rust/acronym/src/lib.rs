pub fn abbreviate(phrase: &str) -> String {
    phrase
        .split(|c| c == '-' || c == ' ')
        .fold(String::new(), |mut acc, w| {
            let chars: Vec<char> = w.replace("_", "").chars().collect();
            let len = chars.len();

            if len > 0 {
                acc.push(chars[0].to_ascii_uppercase());
            }

            if len > 2 {
                for i in 1..chars.len() - 1 {
                    if chars[i].is_lowercase() && chars[i + 1].is_uppercase() {
                        acc.push(chars[i + 1]);
                    }
                }
            }

            acc
        })
}
