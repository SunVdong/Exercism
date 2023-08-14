use std::cmp::Ordering;

pub fn find<R: AsRef<[T]>, T: Ord>(sorted_data: R , key: T) -> Option<usize> {
    let sorted_data = sorted_data.as_ref();
    let mid = sorted_data.len()/2;

    match key.cmp(sorted_data.get(mid)?) {
        Ordering::Equal => Some(mid),
        Ordering::Less => find(&sorted_data[..mid], key),
        Ordering::Greater => find(&sorted_data[mid+1..], key).map(|i| mid+1+i),
    }
}