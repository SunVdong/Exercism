// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut map = HashMap::new();

    for val in magazine.iter() {
        if map.get(val) != None {
            map.insert(val, map.get(val).unwrap() + 1);
        } else {
            map.insert(val, 1);
        }
    }

    for x in note.iter() {
        if !map.contains_key(x) {
            return false;
        }
        if map.get(x).unwrap().to_owned() > 1 {
            map.insert(x, map.get(x).unwrap() - 1);
        } else {
            map.remove(x);
        }
    }
    true
}
