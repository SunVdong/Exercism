use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    if possible_anagrams.is_empty() {
        return HashSet::new();
    }

    let mut h: HashSet<&str> = HashSet::new();
    let mut word_lower = word.to_lowercase().chars().collect::<Vec<char>>();
    word_lower.sort_unstable();

    possible_anagrams.iter().for_each(|&x| {
        if x.len() == word.len() && x.to_lowercase() != word.to_lowercase() {
            let mut x_lower = x.to_lowercase().chars().collect::<Vec<char>>();
            x_lower.sort_unstable();

            if x_lower == word_lower {
                h.insert(x);
            }
        }
    });

    h
}
