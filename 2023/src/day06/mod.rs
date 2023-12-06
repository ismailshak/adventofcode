use regex::Regex;

use crate::puzzle::{Puzzle, PuzzlePart};

fn parse_digits(r: &Regex, line: &str) -> Vec<u32> {
    r.captures_iter(line)
        .map(|c| c.get(1).unwrap().as_str().parse::<u32>().unwrap())
        .collect::<Vec<u32>>()
}

fn part1() -> u32 {
    let digit_regex = Regex::new(r"(\d+)").unwrap();

    let lines = include_str!("input.txt").split("\n").collect::<Vec<&str>>();

    let times = parse_digits(&digit_regex, lines[0]);
    let distances = parse_digits(&digit_regex, lines[1]);

    times
        .iter()
        .enumerate()
        .map(|(i, t)| {
            let mut ways = 0;
            for x in 0..*t {
                let dx = x * (t - x);
                if dx > distances[i] {
                    ways += 1;
                }
            }

            ways
        })
        .fold(1, |acc, n| acc * n)
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 6,
        title: "Wait For It",
        part1: PuzzlePart::new("Product of all different ways to win", part1),
        part2: PuzzlePart::new("N/A", || 0),
    }
}
