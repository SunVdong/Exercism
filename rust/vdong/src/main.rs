use std::collections::HashMap;

fn main() {
    // let magazine = "i've got some lovely coconuts"
    //     .split_whitespace()
    //     .collect::<Vec<&str>>();

    // let mut hash = HashMap::new();
    // for (i, val) in magazine.iter().enumerate() {
    //     hash.insert(val, i);
    // }

    // let note = "I've got som coconuts"
    //     .split_whitespace()
    //     .collect::<Vec<&str>>();

    // for x in note.iter() {
    //     if !hash.contains_key(x) {
    //         return false;
    //     }
    // }
    // true

    let mut map = HashMap::new();
    map.insert("a", 1);
    map.insert("b", 2);
    map.insert("c", 3);

    for (key, val) in map.iter() {
        if map.get(key) != None {
            println!("{}", map.get(key).unwrap() + 1);
        }
        println!("key: {} val: {}", key, val);
    }
}
