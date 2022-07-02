pub fn factors(n: u64) -> Vec<u64> {
    let mut n = n;
    let mut start: u64 = 2;
    let mut v = Vec::new();

    while n != 1 {
        if n % start == 0 {
            v.push(start);
            n = n / start;
        } else {
            start += 1
        }
    }
    v
}
