pub fn is_armstrong_number(num: u32) -> bool {
    let num_str = num.to_string();
    let len = num_str.len();

    num == (0..len)
        .map(|x| (&num_str[x..x+1].parse::<u32>().unwrap()).pow(len as u32))
        .sum()
}
