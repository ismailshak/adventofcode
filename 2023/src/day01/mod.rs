use crate::puzzle::{Puzzle, PuzzlePart};

fn part1() -> u32 {
    include_str!("./input.txt")
        .lines()
        .map(|l| {
            let digits = l
                .chars()
                .filter(|x| x.is_ascii_digit())
                .collect::<Vec<char>>();

            let first = digits.first().expect("Failed to find first");
            let last = digits.last().expect("Failed to find last");

            format!("{}{}", first, last).parse::<u32>().unwrap()
        })
        .sum::<u32>()
}

fn part2() -> u32 {
    include_str!("./input.txt")
        .lines()
        .map(|l| {
            let parsed = l
                .to_string()
                // Wrapping with string equiv. so that overlaps like `sevenine` are handled
                .replace("one", "one1one")
                .replace("two", "two2two")
                .replace("three", "three3three")
                .replace("four", "four4four")
                .replace("five", "five5five")
                .replace("six", "six6six")
                .replace("seven", "seven7seven")
                .replace("eight", "eight8eight")
                .replace("nine", "nine9nine")
                .chars()
                .filter_map(|x| x.to_digit(10))
                .collect::<Vec<u32>>();

            let first = parsed.first().expect("Failed to find first");
            let last = parsed.last().expect("Failed to find last");

            first * 10 + last
        })
        .sum::<u32>()
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 1,
        title: "Trebuchet?!",
        part1: PuzzlePart::new("Sum of digit calibration values", part1),
        part2: PuzzlePart::new("Sum of all calibration values", part2),
    }
}
