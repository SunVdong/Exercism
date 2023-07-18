#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}

///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return an `Err(.)` if the conversion is impossible.
/// The tests do not test for specific values inside the `Err(.)`.
///
///
/// You are allowed to change the function signature as long as all test still pass.
///
///
/// Example:
/// Input
///   number: &[4, 2]
///   from_base: 10
///   to_base: 2
/// Result
///   Ok(vec![1, 0, 1, 0, 1, 0])
///
/// The example corresponds to converting the number 42 from decimal
/// which is equivalent to 101010 in binary.
///
///
/// Notes:
///  * The empty slice ( "[]" ) is equal to the number 0.
///  * Never output leading 0 digits, unless the input number is 0, in which the output must be `[0]`.
///    However, your function must be able to process input with leading 0 digits.
///
pub fn convert(number: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, Error> {
    // 检查输入的进制是否有效
    if from_base < 2 {
        return Err(Error::InvalidInputBase);
    }

    if to_base < 2 {
        return Err(Error::InvalidOutputBase);
    }
     
    // 检查输入数字是否为空
    if number.is_empty() {
        return Ok(vec![0]);
    }

    // 检查输入数字中是否包含超出进制范围的值
    let max_digit = number.iter().max().copied().unwrap_or(0);
    if max_digit >= from_base {
        return Err(Error::InvalidDigit(max_digit));
    }

    // 将输入数字从 from_base 转换为十进制
    let mut decimal_number: u32 = 0;
    let mut power: u32 = 1;
    for &digit in number.iter().rev() {
        decimal_number += digit * power;
        power *= from_base;
    }

    // 将十进制数字转换为 to_base
    let mut converted_number = Vec::new();
    if decimal_number == 0 {
        converted_number.push(0);
    }
    while decimal_number > 0 {
        let digit = decimal_number % to_base;
        converted_number.push(digit);
        decimal_number /= to_base;
    }

    Ok(converted_number.into_iter().rev().collect())
}
