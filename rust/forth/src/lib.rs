use std::collections::BTreeMap;

pub type Value = i32;
pub type Result = std::result::Result<(), Error>;

#[derive(Debug)]
pub struct Forth {
    data: Vec<Value>,
    syntax_tree: BTreeMap<String, String>,
}

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    DivisionByZero,
    StackUnderflow,
    UnknownWord,
    InvalidWord,
}

impl Forth {
    pub fn new() -> Forth {
        Forth {
            data: Vec::new(),
            syntax_tree: BTreeMap::new(),
        }
    }

    pub fn stack(&self) -> &[Value] {
        &self.data
    }

    pub fn eval(&mut self, input: &str) -> Result {
        let mut status: u8 = 0;
        let mut word_name = String::new();
        let mut definition = String::new();

        for item in input.split_whitespace() {
            if let Ok(num) = item.parse::<Value>() {
                self.data.push(num);
            } else {
                match (item, status, self.syntax_tree.get("dup-twice")) {
                    // ("_", 0, Some(str)) => {
                    //     let _ = self.eval(&str);
                    // }
                    ("+", 0, _) => match (self.data.pop(), self.data.pop()) {
                        (Some(r), Some(l)) => self.data.push(l + r),
                        _ => return Err(Error::StackUnderflow),
                    },
                    ("-", 0, _) => match (self.data.pop(), self.data.pop()) {
                        (Some(r), Some(l)) => self.data.push(l - r),
                        _ => return Err(Error::StackUnderflow),
                    },
                    ("*", 0, _) => match (self.data.pop(), self.data.pop()) {
                        (Some(r), Some(l)) => self.data.push(l * r),
                        _ => return Err(Error::StackUnderflow),
                    },
                    ("/", 0, _) => match (self.data.pop(), self.data.pop()) {
                        (Some(0), _) => return Err(Error::DivisionByZero),
                        (Some(r), Some(l)) => self.data.push(l / r),
                        _ => return Err(Error::StackUnderflow),
                    },
                    (s, 0, _) if s.to_lowercase() == "dup" => match self.data.pop() {
                        Some(v) => {
                            self.data.push(v);
                            self.data.push(v);
                        }
                        _ => return Err(Error::StackUnderflow),
                    },
                    (s, 0, _) if s.to_lowercase() == "drop" => match self.data.pop() {
                        Some(_) => (),
                        _ => return Err(Error::StackUnderflow),
                    },
                    (s, 0, _) if s.to_lowercase() == "swap" => {
                        match (self.data.pop(), self.data.pop()) {
                            (Some(r), Some(l)) => {
                                self.data.push(r);
                                self.data.push(l);
                            }
                            _ => return Err(Error::StackUnderflow),
                        }
                    }
                    (s, 0, _) if s.to_lowercase() == "over" => {
                        match (self.data.pop(), self.data.pop()) {
                            (Some(r), Some(l)) => {
                                self.data.push(l);
                                self.data.push(r);
                                self.data.push(l);
                            }
                            _ => return Err(Error::StackUnderflow),
                        }
                    }
                    (":", 0, _) => {
                        status = 1;
                    }
                    (s, 1, _) => {
                        word_name = s.to_string();
                        status = 2;
                    }
                    (s, 2, _) if s != ";" => {
                        definition = definition + " " + s;
                    }
                    (";", 2, _) => {
                        self.syntax_tree.insert(word_name, definition);
                        word_name = String::new();
                        definition = String::new();
                        status = 0;
                    }
                    _ => return Err(Error::UnknownWord),
                }
            }
        }

        //         if let Some(&ref v) = self.syntax_tree.get("dup-twice") {
        //     let _ = self.eval(&("1".to_owned() + &v.clone()));
        // }

        Ok(())
    }
}
