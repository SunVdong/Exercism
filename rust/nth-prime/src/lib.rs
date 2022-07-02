pub fn nth(n: u32) -> u32 {
    let mut index = 1;

    let mut i: u32 = 2;
    while index < n {
        if is_prime(i) {
            index += 1;
        }
        i = i + 1;
    }

    i
}

pub fn is_prime(num: u32) -> bool {
    if num <= 3 {
        return num > 1;
    }

    if num % 6 != 1 && num % 6 != 5 {
        return false;
    }

    let sqrt = (num as f64).sqrt() as u32 + 1;
    let mut i: u32 = 5;
    while i < sqrt {
        if num % i == 0 || num % (i + 2) == 0 {
            return false;
        }
        i += 5;
    }

    true
}
