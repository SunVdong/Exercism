// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]
const CAR_PER_HOUR:u32 = 221;

pub fn production_rate_per_hour(speed: u8) -> f64 {
   (speed as f64) * (CAR_PER_HOUR as f64) * match speed {
        0 => 0.00,
        1..=4 => 1.0,
        5..=8 => 0.9,
        9..=10 => 0.77,
        _ => panic!("error"),
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) as u32) / 60
}
