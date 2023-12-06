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

fn parse_digit(r: &Regex, line: &str) -> u64 {
    r.captures_iter(line)
        .map(|c| c.get(1).unwrap().as_str())
        .collect::<Vec<&str>>()
        .join("")
        .parse()
        .unwrap()
}

fn part2() -> u32 {
    let digit_regex = Regex::new(r"(\d+)").unwrap();
    let lines = include_str!("input.txt").split("\n").collect::<Vec<&str>>();
    let time = parse_digit(&digit_regex, lines[0]);
    let distance = parse_digit(&digit_regex, lines[1]);

    // Not sure why this doesn't work... It's super fast tho lol
    //
    // max_hold_to_win(time, distance) - min_hold_to_win(time, distance) + 1

    let mut ways = 0;
    for x in 0..time {
        let dx = x * (time - x);
        if dx > distance {
            ways += 1;
        }
    }

    ways
}

pub fn get<'a>() -> Puzzle<'a, u32> {
    Puzzle {
        day: 6,
        title: "Wait For It",
        part1: PuzzlePart::new("Product of all different ways to win", part1),
        part2: PuzzlePart::new("Number of ways to win the long race", part2),
    }
}

// Moving my failed attempt at using my math skills to solve this problem down here
// ================================================================================

// Solution for part 1 works fine, but part 2 uses crazy large numbers...
//
// Math notes:
//
// t = max time
// d = distance to beat
// x = time to hold the button
//
// To get how far the boat will travel for a given time x, we use:
// d = x * (t - x)
//
// Rearranging that formula, you get:
// x^2 - tx + d = 0 (quadratic formula)
//
// The quadratic formula is:
// x = (-b +- sqrt(b^2 - 4ac)) / 2a
//
// In our case, a = 1, b = -t, c = d
// x = (t +- sqrt(t^2 - 4d)) / 2
//
// Where +- is the two solutions to the quadratic formula. In our
// case, the 2 solutions for x are the minimum and maximum time
// to hold the button and still win the race.
//
// To get the minimum time to hold the button, we use the negative
// To get the maximum time to hold the button, we use the positive
//
// Crude diagram to help visualize: (this is supposed to be a quadratic curve)
//
// |
// |
// |*              *
// | *            *
// |  *          *
// |---*--------*----|
// 0    *      *     t
// |      *   *
// |        *
// |
//
// The two points where the curve crosses the x axis are the min and max button hold times. So we
// gotta solve for x.

#[allow(dead_code)]
fn min_hold_to_win(t: u64, d: u64) -> u32 {
    let t = t as f64;
    let d = d as f64;
    let x = ((t * t + 4.0 * d).sqrt() - t) / 2.0;
    x.ceil() as u32
}

#[allow(dead_code)]
fn max_hold_to_win(t: u64, d: u64) -> u32 {
    let t = t as f64;
    let d = d as f64;
    let x = ((t * t - 4.0 * d).sqrt() + t) / 2.0;
    x.floor() as u32
}
