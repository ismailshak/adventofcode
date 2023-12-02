pub struct PuzzlePart<'a, T>
where
    T: std::fmt::Debug,
    T: std::fmt::Display,
{
    pub message: &'a str,
    pub solve: fn() -> T,
}

impl<'a, T> PuzzlePart<'a, T>
where
    T: std::fmt::Debug,
    T: std::fmt::Display,
{
    pub fn new(message: &'a str, solve: fn() -> T) -> Self {
        PuzzlePart { message, solve }
    }
}

// Second type parameter `S` is if part 2 wants to return
// a different data type for some reason
pub struct Puzzle<'a, T, S = T>
where
    T: std::fmt::Debug,
    T: std::fmt::Display,
    S: std::fmt::Debug,
    S: std::fmt::Display,
{
    pub day: u8,
    pub title: &'a str,
    pub part1: PuzzlePart<'a, T>,
    pub part2: PuzzlePart<'a, S>,
}
