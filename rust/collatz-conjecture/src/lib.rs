pub fn collatz(mut n: u64) -> Option<u64> {
    for i in 0.. {
        match n {
            0 => break,
            1 => return Some(i),
            even if even & 1 == 0 => n >>= 1,
            _ => n = n.checked_mul(3)?.checked_add(1)?,
        }
    }
    None
}
