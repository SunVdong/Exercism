#[derive(Debug)]
pub struct ChessPosition {
    rank: i8,
    file: i8,
}

#[derive(Debug)]
pub struct Queen {
    position: ChessPosition,
}

impl ChessPosition {
    pub fn new(rank: i32, file: i32) -> Option<Self> {
        if rank < 0 || rank >= 8 || file < 0 || file >= 8 {
            return None;
        }

        Some(ChessPosition {
            rank: rank as i8,
            file: file as i8,
        })
    }
}

impl Queen {
    pub fn new(position: ChessPosition) -> Self {
        Self { position: position }
    }

    pub fn can_attack(&self, other: &Queen) -> bool {
        self.position.rank == other.position.rank
            || self.position.file == other.position.file
            || (self.position.rank - other.position.rank).abs() == (self.position.file - other.position.file).abs()
    }
}
