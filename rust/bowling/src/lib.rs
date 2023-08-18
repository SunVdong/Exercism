#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

#[derive(Debug, PartialEq, Clone, Copy)]
pub enum ResultType {
    OpenFrame,
    Spare,
    Strike,
    NotFinish,
}

#[derive(Debug)]
pub struct BowlingGame {
    frame_total: usize,
    rolls: Vec<(u16, u16, ResultType)>,
    fill_ball: Vec<u16>,
    current_frame: usize,
}

impl BowlingGame {
    pub fn new() -> Self {
        BowlingGame {
            frame_total: 10,
            rolls: Vec::new(),
            fill_ball: Vec::new(),
            current_frame: 0,
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if pins > 10 {
            return Err(Error::NotEnoughPinsLeft);
        }

        match self.current_frame {
            n if n < self.frame_total => {
                // println!("正常业务流程");
                if let Some(element) = self.rolls.get_mut(self.current_frame) {
                    if element.2 == ResultType::NotFinish {
                        match element.0 + pins {
                            0..=9 => {
                                element.2 = ResultType::OpenFrame;
                                element.1 = pins;
                                self.rolls[self.current_frame] = *element;
                                self.current_frame += 1;
                            }
                            10 => {
                                element.2 = ResultType::Spare;
                                element.1 = pins;
                                self.rolls[self.current_frame] = *element;
                                self.current_frame += 1;
                            }
                            _ => return Err(Error::NotEnoughPinsLeft),
                        }
                    }
                } else {
                    match pins {
                        10 => {
                            self.rolls.push((10, 0, ResultType::Strike));
                            self.current_frame += 1;
                        }
                        _ => self.rolls.push((pins, 0, ResultType::NotFinish)),
                    }
                }
            }
            _ => {
                // println!("处理补球流程");
                match self.rolls[self.frame_total - 1].2 {
                    ResultType::Strike => {
                        if self.fill_ball.len() >= 2 {
                            return Err(Error::GameComplete);
                        }
                        if self.fill_ball.len() == 1
                            && self.fill_ball[0] != 10
                            && self.fill_ball[0] + pins > 10
                        {
                            return Err(Error::NotEnoughPinsLeft);
                        }
                        self.fill_ball.push(pins);
                    }
                    ResultType::Spare => {
                        if self.fill_ball.len() >= 1 {
                            return Err(Error::GameComplete);
                        }
                        self.fill_ball.push(pins);
                    }
                    _ => {
                        return Err(Error::GameComplete);
                    }
                }
            }
        };

        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        println!("{:?}", self);

        if self.current_frame != self.frame_total {
            return None;
        }
        if self.rolls[self.frame_total - 1].2 == ResultType::Strike && self.fill_ball.len() < 2
            || self.rolls[self.frame_total - 1].2 == ResultType::Spare && self.fill_ball.len() < 1
        {
            return None;
        }

        let mut sum = 0;
        for (idx, (a, b, res)) in self.rolls.iter().enumerate() {
            if idx < 9 {
                match res {
                    ResultType::OpenFrame => {
                        sum += a + b;
                    }
                    ResultType::Strike => {
                        if self.rolls[idx + 1].2 == ResultType::Strike {
                            if let Some(e) = self.rolls.get(idx + 2) {
                                sum = sum + 20 + e.0
                            } else {
                                sum = sum + 20 + self.fill_ball[0]
                            }
                        } else {
                            sum += 10 + self.rolls[idx + 1].0 + self.rolls[idx + 1].1;
                        }
                    }
                    ResultType::Spare => {
                        sum += 10 + self.rolls[idx + 1].0;
                    }
                    ResultType::NotFinish => return None,
                }
            }
        }
        sum = sum
            + self.rolls[self.frame_total - 1].0
            + self.rolls[self.frame_total - 1].1
            + self.fill_ball.iter().sum::<u16>();
        return Some(sum);
    }
}
