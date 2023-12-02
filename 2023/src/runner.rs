use crate::day01;
use crate::day02;
use crate::puzzle::Puzzle;
use colored::Colorize;

// Wrapper enum for the puzzles return type combinations
enum PuzzleResult<'a> {
    Int(Puzzle<'a, u32>),
}

fn print_header(day: u8, title: &str) {
    let header = format!(
        "🎄 {} {}: {} 🎄",
        "Day".bold(),
        day.to_string().bold(),
        title.bold().magenta(),
    );
    println!();
    println!("{}", header);
    println!("{}", "-".repeat(day.to_string().len() + title.len() + 12));
}

fn print_result<T>(result: T, message: &str, duration: u128)
where
    T: std::fmt::Display,
{
    println!(
        "{}: {} {}",
        message,
        result.to_string().green(),
        format!("({:.3}ms)", (duration as f64 / 1_000_000.0)).yellow()
    );
}

fn solve_puzzle<T>(puzzle: &Puzzle<T>)
where
    T: std::fmt::Debug,
    T: std::fmt::Display,
{
    print_header(puzzle.day, puzzle.title);

    let start = std::time::Instant::now();
    let result = (puzzle.part1.solve)();
    let duration = start.elapsed().as_nanos();
    print_result(result, puzzle.part1.message, duration);

    println!();

    let start = std::time::Instant::now();
    let result = (puzzle.part2.solve)();
    let duration = start.elapsed().as_nanos();
    print_result(result, puzzle.part2.message, duration);
}

pub fn run(day: &str) {
    let puzzle = match day {
        "1" | "01" => Some(PuzzleResult::Int(day01::get())),
        "2" | "02" => Some(PuzzleResult::Int(day02::get())),
        _ => None,
    };

    match puzzle {
        Some(PuzzleResult::Int(p)) => solve_puzzle(&p),
        None => {
            eprintln!("Day {} has not been implemented", day);
            std::process::exit(1);
        }
    }
}
