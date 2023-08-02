const STUDENTS:[&str;12] = [
        "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph",
        "Kincaid", "Larry",
    ];

pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    let position = STUDENTS.iter().position(|&x| x == student).unwrap()*2;

    let plant_from_char = |x| match x {
        'G' => "grass",
        'C' => "clover",
        'R' => "radishes",
        'V' => "violets",
        _ => "",
    };

    diagram.lines().flat_map(|line|{
        line[position..=position+1].chars().map(plant_from_char)
    }).collect()
}
