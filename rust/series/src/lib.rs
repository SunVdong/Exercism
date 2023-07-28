pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut res: Vec<String> = vec![];
    let l = digits.len();

    for i in 0..l {
        if i+len <= l {
            res.push(digits.get(i..i+len).unwrap().to_string());
        }
    }
    
    return res;
}
