pub fn number(user_number: &str) -> Option<String> {
    let digits: String = user_number.chars().filter(|c| c.is_digit(10)).collect();

    match (
        digits.len(),
        digits.chars().nth(0).unwrap(),
        digits.chars().nth(1).unwrap(),
        digits.chars().nth(3).unwrap(),
        digits.chars().nth(4).unwrap(),
    ) {
        (10, '2'..='9', _, '2'..='9', _) => Some(digits),
        (11, '1', '2'..='9', _, '2'..='9') => Some(digits[1..].to_string()),
        _ => None,
    }
}
