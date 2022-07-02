use regex::Regex;

pub fn brackets_are_balanced(string: &str) -> bool {
    let re = Regex::new(r"[^\{\[\(\)\]\}]").unwrap();
    let mut str = re.replace_all(string, "").to_string();

    let result = loop {
        let res = str.replace("()", "").replace("{}", "").replace("[]", "");
        if res == str {
            break res;
        }

        str = res;
    };

    result.is_empty()
}
