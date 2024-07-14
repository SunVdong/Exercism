use forth::Forth;

fn main() {
    let mut f = Forth::new();
    let _ = f.eval(": dup-twice dup dup ;");

    println!("{:?}", f.syntax_tree.get("dup-twice"));

    // assert!(f.eval("1 dup-twice").is_ok());
    // assert_eq!(vec![1, 1, 1], f.stack());
}
