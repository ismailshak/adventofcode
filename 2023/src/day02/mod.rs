use crate::puzzle::{Puzzle, PuzzlePart};
use regex::Regex;

/// Extracts the ID value for a "game"
fn get_game_id(r: &Regex, game: &str) -> u32 {
    let id_caps = r.captures(game).unwrap();
    id_caps.get(1).unwrap().as_str().parse().unwrap()
}

/// Extracts the digit for each capture.
/// (e.g. in "2 red", a parsed 2 would be returned)
///
/// If no matches foud for a color, a 0 is returned
fn extract_count(c: &regex::Captures<'_>) -> u32 {
    match c.get(1) {
        Some(v) => v.as_str().parse::<u32>().unwrap(),
        None => 0,
    }
}

fn part1() -> u32 {
    let id_regex = Regex::new(r"Game (\d+)").unwrap();
    let red_regex = Regex::new(r"(\d+) red").unwrap();
    let green_regex = Regex::new(r"(\d+) green").unwrap();
    let blue_regex = Regex::new(r"(\d+) blue").unwrap();

    include_str!("./input.txt")
        .lines()
        .filter_map(|l| {
            let reds = red_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0); // Handles if iterator has no elements when max() is called

            if reds > 12 {
                return None;
            }

            let greens = green_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0);

            if greens > 13 {
                return None;
            }

            let blues = blue_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0);

            if blues > 14 {
                return None;
            }

            Some(get_game_id(&id_regex, l))
        })
        .sum::<u32>()
}

fn part2() -> u32 {
    let red_regex = Regex::new(r"(\d+) red").unwrap();
    let green_regex = Regex::new(r"(\d+) green").unwrap();
    let blue_regex = Regex::new(r"(\d+) blue").unwrap();

    include_str!("./input.txt")
        .lines()
        .map(|l| {
            let reds = red_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0);

            let greens = green_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0);

            let blues = blue_regex
                .captures_iter(l)
                .map(|c| extract_count(&c))
                .max()
                .unwrap_or(0);

            reds * greens * blues
        })
        .sum::<u32>()
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 2,
        title: "Cube Conundrum",
        part1: PuzzlePart::new("Sum of valid games", part1),
        part2: PuzzlePart::new("Sum of minimum cubes", part2),
    }
}
