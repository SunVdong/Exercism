use enum_iterator::IntoEnumIterator;
use int_enum::{IntEnum, IntEnumError};

#[repr(usize)]
#[derive(Debug, PartialEq, Copy, Clone, IntEnum, IntoEnumIterator)]
pub enum ResistorColor {
    Black = 0,
    Brown = 1,
    Red = 2,
    Orange = 3,
    Yellow = 4,
    Green = 5,
    Blue = 6,
    Violet = 7,
    Grey = 8,
    White = 9,
}

pub fn color_to_value(_color: ResistorColor) -> usize {
    // _color as usize
    _color.int_value()
}

pub fn value_to_color_string(value: usize) -> String {
    let cast: Result<ResistorColor, IntEnumError<ResistorColor>> = ResistorColor::from_int(value);
    match cast {
        Ok(x) => format!("{:?}", x),
        Err(_) => String::from("value out of range"),
    }
}

pub fn colors() -> Vec<ResistorColor> {
    // Vec::from_iter(ResistorColor::into_enum_iter())

    ResistorColor::into_enum_iter().collect()
}
