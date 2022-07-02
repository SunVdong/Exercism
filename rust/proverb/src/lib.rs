pub fn build_proverb(list: &[&str]) -> String {
    if list.is_empty() {
        return String::new();
    }
    let length = list.len();
    let res = (0..length - 1)
        .map(|x| format!("For want of a {} the {} was lost.\n", list[x], list[x + 1]))
        .collect::<Vec<_>>()
        .join("");

    res + &format!("And all for the want of a {}.", list[0])
}