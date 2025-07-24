pub struct CircularBuffer<T> {
    buffer: Vec<Option<T>>,
    capacity: usize,
    head: usize,
    tail: usize,
    size: usize,
}

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    EmptyBuffer,
    FullBuffer,
}

impl<T: Clone> CircularBuffer<T> {
    pub fn new(capacity: usize) -> Self {
        if capacity == 0 {
            panic!("Capacity must be greater than zero.");
        }
        let mut buffer = Vec::with_capacity(capacity);
        for _ in 0..capacity {
            buffer.push(None);
        }
        CircularBuffer {
            buffer,
            capacity,
            head: 0,
            tail: 0,
            size: 0,
        }
    }

    pub fn write(&mut self, element: T) -> Result<(), Error> {
        if self.size == self.capacity {
            return Err(Error::FullBuffer);
        }
        self.buffer[self.head] = Some(element);
        self.head = (self.head + 1) % self.capacity;
        self.size += 1;
        Ok(())
    }

    pub fn read(&mut self) -> Result<T, Error> {
        if self.size == 0 {
            return Err(Error::EmptyBuffer);
        }
        let element = self.buffer[self.tail].take().unwrap();
        self.tail = (self.tail + 1) % self.capacity;
        self.size -= 1;
        Ok(element)
    }

    pub fn clear(&mut self) {
        for i in 0..self.capacity {
            self.buffer[i] = None;
        }
        self.head = 0;
        self.tail = 0;
        self.size = 0;
    }

    pub fn overwrite(&mut self, element: T) {
        if self.size == self.capacity {
            self.tail = (self.tail + 1) % self.capacity; // Overwrite the oldest element
        } else {
            self.size += 1;
        }
        self.buffer[self.head] = Some(element);
        self.head = (self.head + 1) % self.capacity;
    }
}
