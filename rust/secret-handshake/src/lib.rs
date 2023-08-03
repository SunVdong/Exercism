const SECRET: [&str; 4] = ["wink", "double blink", "close your eyes", "jump"];

pub fn actions(n: u8) -> Vec<&'static str> {
    let mut res = Vec::new();
    let mut offset_list: [u8; 4] = [0, 1, 2, 3];

    if n & 16 == 16 {
        offset_list.reverse();
    }
    
    for &v in offset_list.iter() {
        if (1 << v) & n == (1 << v) {
            res.push(SECRET[v as usize]);
        }
    }

    res
}
