use std::ops::Add;

pub struct Triangle<T> {
    a: T,
    b: T,
    c: T,
}

impl<T: PartialOrd + Copy + Add> Triangle<T> {
    pub fn build(arr: [T; 3]) -> Option<Self>
    where
        <T as Add>::Output: PartialOrd<T>,
    {
        let a = arr[0];
        let b = arr[1];
        let c = arr[2];

        if a + b > c && a + c > b && b + c > a {
            Some(Self { a, b, c })
        } else {
            None
        }
    }

    pub fn is_equilateral(&self) -> bool {
        self.a == self.b && self.b == self.c
    }

    pub fn is_isosceles(&self) -> bool {
        self.a == self.b || self.b == self.c || self.a == self.c
    }

    pub fn is_scalene(&self) -> bool {
        self.a != self.b && self.b != self.c && self.a != self.c
    }
}
