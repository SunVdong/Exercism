// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug)]
pub struct Duration(f64);

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Self((s as f64) / 31557600.00)
    }
}

pub trait Planet {
    fn period() -> f64;
    fn years_during(d: &Duration) -> f64 {
        d.0 / Self::period()
    }
}

// macro_rules! planet {
//     ($n:ident, $p:expr) => {
//         pub struct $n; impl Planet for $n { fn period() -> f64 { $p } }
//     }
// }

// planet!(Earth, 1.0);
// planet!(Mercury, 0.2408467);
// planet!(Venus, 0.61519726);
// planet!(Mars, 1.8808158);
// planet!(Jupiter, 11.862615);
// planet!(Saturn, 29.447498);
// planet!(Uranus, 84.016846);
// planet!(Neptune, 164.79132);

macro_rules! boilerplate {
    ($($t:ident => $e:expr),+) => {
        $(
            pub struct $t;
            impl Planet for $t {
                fn period() -> f64 {
                    $e
                }
            }
        )*
    }
}

boilerplate!(
    Mercury => 0.2408467,
    Venus => 0.61519726,
    Earth => 1.0,
    Mars => 1.8808158,
    Jupiter => 11.862615,
    Saturn => 29.447498,
    Uranus => 84.016846,
    Neptune => 164.7913
);
