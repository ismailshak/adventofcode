use crate::puzzle::{Puzzle, PuzzlePart};
use regex::Regex;

fn parse_digits(r: &Regex, line: &str) -> Vec<u32> {
    r.captures_iter(line)
        .map(|c| c.get(1).unwrap().as_str().parse::<u32>().unwrap())
        .collect::<Vec<u32>>()
}

fn apply_transformations(
    list: &mut Vec<u32>,
    state: &mut Vec<bool>,
    destination_start: u32,
    source_start: u32,
    range: u32,
) {
    let diff = (destination_start as i32) - (source_start as i32);
    let range_end = source_start + range;
    for i in 0..list.len() {
        let curr = list[i];
        if (source_start..range_end).contains(&curr) && !state[i] {
            list[i] = ((curr as i32) + diff) as u32;
            state[i] = true;
        }
    }
}

fn part1() -> u32 {
    let digit_regex = Regex::new(r"(\d+)").unwrap();
    let mut values = Vec::new();
    let mut category_state = Vec::new();

    include_str!("input.txt")
        .lines()
        .enumerate()
        .for_each(|(i, l)| {
            if i == 0 {
                // Initialize transform with initial set of seed numbers
                values = parse_digits(&digit_regex, l);
                category_state = vec![false; values.len()];
                return;
            }

            // New lines mark the end of a mapping category, so we reset statuses
            if l.is_empty() {
                category_state = vec![false; values.len()];
                return;
            }

            // If no digits, then this is a mapping header line
            if !digit_regex.is_match(l) {
                return;
            }

            // Otherwise, we compute the mapping ranges
            let mapping = parse_digits(&digit_regex, l);

            apply_transformations(
                &mut values,
                &mut category_state,
                mapping[0],
                mapping[1],
                mapping[2],
            );
        });

    *values.iter().min().unwrap()
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 5,
        title: "If You Give A Seed A Fertilizer",
        part1: PuzzlePart::new("Lowest location number", part1),
        part2: PuzzlePart::new("N/A", || 0),
    }
}
