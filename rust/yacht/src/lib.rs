use std::collections::HashMap;

pub enum Category {
    Ones,
    Twos,
    Threes,
    Fours,
    Fives,
    Sixes,
    FullHouse,
    FourOfAKind,
    LittleStraight,
    BigStraight,
    Choice,
    Yacht,
}

type Dice = [u8; 5];
pub fn score(dice: Dice, category: Category) -> u8 {
    let me = |(acc, target), &x| (if x == target { acc + target } else { acc }, target);

    match category {
        Category::Ones => dice.iter().fold((0_u8, 1_u8), me).0,
        Category::Twos => dice.iter().fold((0_u8, 2_u8), me).0,
        Category::Threes => dice.iter().fold((0_u8, 3_u8), me).0,
        Category::Fours => dice.iter().fold((0_u8, 4_u8), me).0,
        Category::Fives => dice.iter().fold((0_u8, 5_u8), me).0,
        Category::Sixes => dice.iter().fold((0_u8, 6_u8), me).0,
        Category::FullHouse => {
            let mut count_map = HashMap::new();
            for &x in &dice {
                let count = count_map.entry(x).or_insert(0);
                *count += 1;
            }

            let mut found_3 = false;
            let mut found_2 = false;
            let mut key_3 = 0;
            let mut key_2 = 0;

            for (key, count) in count_map {
                if count == 3 {
                    found_3 = true;
                    key_3 = key;
                } else if count == 2 {
                    found_2 = true;
                    key_2 = key;
                } else {
                    return 0;
                }
            }
            if found_3 && found_2 {
                key_2 * 2 + key_3 * 3
            } else {
                0
            }
        }
        Category::FourOfAKind => {
            let mut count_map = HashMap::new();
            for &x in &dice {
                let count = count_map.entry(x).or_insert(0);
                *count += 1;
            }
            for (key, count) in count_map {
                if count >= 4 {
                    return  key*4;
                } 
            }
            0
        },
        Category::LittleStraight => {
            if dice.iter().fold(0, |acc, &x| x ^ acc) == (1 ^ 2 ^ 3 ^ 4 ^ 5) {
                30
            } else {
                0
            }
        }
        Category::BigStraight => {
            if dice.iter().fold(0, |acc, &x| x ^ acc) == (2 ^ 3 ^ 4 ^ 5 ^ 6) {
                30
            } else {
                0
            }
        }
        Category::Choice => dice.iter().sum(),
        Category::Yacht => {
            if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4]
            {
                50
            } else {
                0
            }
        }
    }
}
