use crate::puzzle::{Puzzle, PuzzlePart};
use regex::Regex;
use std::{collections::HashMap, u32};

fn extract_winning_numbers<'a>(r: &Regex, line: &'a str) -> HashMap<&'a str, u32> {
    let mut winning_map = HashMap::new();

    r.captures_iter(line).for_each(|c| {
        c.get(1).unwrap().as_str().trim().split(" ").for_each(|n| {
            // A single digit is right aligned, so an empty string shows up
            if n.is_empty() {
                return;
            }

            winning_map.insert(n, 0);
        });
    });

    winning_map
}

fn find_win_count<'a>(r: &Regex, line: &'a str, winnings: &mut HashMap<&'a str, u32>) -> u32 {
    r.captures_iter(line)
        .map(|c| {
            c.get(1)
                .unwrap()
                .as_str()
                .trim()
                .split(" ")
                .map(|n| match winnings.get(n) {
                    Some(v) => {
                        // Only count winnings one time if we have dupes on the right
                        if *v != 0 {
                            return 0;
                        }

                        winnings.insert(n, 1).unwrap() + 1
                    }
                    None => 0,
                })
                .sum::<u32>()
        })
        .sum::<u32>()
}

fn part1() -> u32 {
    let winners_regex = Regex::new(r":(.*?)\|").unwrap();
    let draws_regex = Regex::new(r"\|(.*?)$").unwrap();
    include_str!("input.txt")
        .lines()
        .map(|l| {
            let mut winnings = extract_winning_numbers(&winners_regex, l);
            let count = find_win_count(&draws_regex, l, &mut winnings);

            if count == 0 {
                return 0;
            }

            0 | 1 << count - 1
        })
        .sum::<u32>()
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 4,
        title: "Scratchcards",
        part1: PuzzlePart::new("Total winning points", part1),
        part2: PuzzlePart::new("N/A", || 0),
    }
}
