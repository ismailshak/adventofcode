use crate::puzzle::{Puzzle, PuzzlePart};
use regex::Regex;

fn is_symbol(c: &char) -> bool {
    c.is_ascii_punctuation() && (*c) != '.'
}

fn is_adjacent(location: usize, start: usize, end: usize) -> bool {
    let location = location as i32;
    let start = start as i32;
    let end = (end as i32) - 1;

    (location >= start && location <= end) || (location + 1 == start) || (location - 1 == end)
}

fn sum_part_numbers(r: &Regex, location: usize, line: &str) -> u32 {
    r.find_iter(line)
        .filter_map(|m| {
            if !is_adjacent(location, m.start(), m.end()) {
                return None;
            }

            m.as_str().parse::<u32>().ok()
        })
        .sum::<u32>()
}

fn part1() -> u32 {
    let lines: Vec<&str> = include_str!("input.txt").lines().collect();
    let num_regex = Regex::new(r"\d+").unwrap();

    lines
        .windows(3)
        .map(|w| {
            w.get(1)
                .unwrap()
                .chars()
                .enumerate()
                .map(|(i, c)| {
                    if !is_symbol(&c) {
                        return 0;
                    }

                    let top_sum = sum_part_numbers(&num_regex, i, w.get(0).unwrap());
                    let mid_sum = sum_part_numbers(&num_regex, i, w.get(1).unwrap());
                    let bottom_sum = sum_part_numbers(&num_regex, i, w.get(2).unwrap());

                    top_sum + mid_sum + bottom_sum
                })
                .sum::<u32>()
        })
        .sum::<u32>()
}

fn find_gear_numbers(r: &Regex, location: usize, line: &str) -> Vec<u32> {
    r.find_iter(line)
        .filter_map(|m| {
            if !is_adjacent(location, m.start(), m.end()) {
                return None;
            }

            m.as_str().parse::<u32>().ok()
        })
        .collect::<Vec<u32>>()
}

fn part2() -> u32 {
    let lines: Vec<&str> = include_str!("input.txt").lines().collect();
    let num_regex = Regex::new(r"\d+").unwrap();

    lines
        .windows(3)
        .map(|w| {
            w.get(1)
                .unwrap()
                .chars()
                .enumerate()
                .filter_map(|(i, c)| {
                    if c != '*' {
                        return None;
                    }

                    let top = find_gear_numbers(&num_regex, i, w.get(0).unwrap());
                    let mid = find_gear_numbers(&num_regex, i, w.get(1).unwrap());
                    let bottom = find_gear_numbers(&num_regex, i, w.get(2).unwrap());

                    let non_zero_values: Vec<u32> = vec![top, mid, bottom]
                        .into_iter()
                        .flatten()
                        .filter(|x| *x != 0)
                        .collect();

                    if non_zero_values.len() != 2 {
                        return None;
                    }

                    Some(non_zero_values.iter().fold(1, |acc, x| acc * x))
                })
                .sum::<u32>()
        })
        .sum::<u32>()
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 3,
        title: "Gear Ratios",
        part1: PuzzlePart::new("Sum of all part numbers", part1),
        part2: PuzzlePart::new("Sum of all gear numbers", part2),
    }
}
