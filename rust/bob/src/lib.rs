pub fn reply(message: &str) -> &str {
    match (
        is_nothing(&message),
        is_question(&message),
        is_yell(&message),
    ) {
        (true, _, _) => "Fine. Be that way!",
        (false, true, true) => "Calm down, I know what I'm doing!",
        (false, true, false) => "Sure.",
        (false, false, true) => "Whoa, chill out!",
        (_, _, _) => "Whatever.",
    }
}

pub fn is_question(message: &str) -> bool {
    message.trim().ends_with("?")
}

pub fn is_yell(message: &str) -> bool {
    message.chars().any(|x| x.is_ascii_alphabetic()) && message.chars().all(|x| !x.is_ascii_alphabetic() || x.is_ascii_uppercase())
}

pub fn is_nothing(message: &str) -> bool {
    message.chars().all(|x| x.is_whitespace())
}
