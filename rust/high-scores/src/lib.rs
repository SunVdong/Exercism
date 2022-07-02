#[derive(Debug)]
pub struct HighScores {
    scores: Vec<u32>,
}

impl HighScores {
    pub fn new(scores: &[u32]) -> Self {
        HighScores {
            scores: scores.to_vec(),
        }
    }

    pub fn scores(&self) -> &[u32] {
        &self.scores
    }

    pub fn latest(&self) -> Option<u32> {
        self.scores.last()
    }

    pub fn personal_best(&self) -> Option<u32> {
        self.scores.max()
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let len = self.scores.len();
        let mid = if len <= 3 { 0 } else { len - 3 };
        let (left, right) = self.scores.sort().split_at(len);
        right
    }
}
