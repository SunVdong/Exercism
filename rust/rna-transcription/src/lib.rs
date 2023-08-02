#[derive(Debug, PartialEq, Eq)]
pub struct Dna(String);

#[derive(Debug, PartialEq, Eq)]
pub struct Rna(String);

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        for (index, nucleotide) in dna.chars().enumerate() {
            if !matches!(nucleotide, 'A' | 'C' | 'G' | 'T') {
                return Err(index);
            }
        }
        
        Ok(Dna(dna.to_string()))
    }

    pub fn into_rna(self) -> Rna {
        let rna_sequence = self.0.chars().map(|nucleotide| {
            match nucleotide {
                'G' => 'C',
                'C' => 'G',
                'T' => 'A',
                'A' => 'U',
                _ => panic!("Invalid nucleotide"),
            }
        }).collect();
        
        Rna(rna_sequence)
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        for (index, nucleotide) in rna.chars().enumerate() {
            if !matches!(nucleotide, 'A' | 'C' | 'G' | 'U') {
                return Err(index);
            }
        }
        
        Ok(Rna(rna.to_string()))
    }
}