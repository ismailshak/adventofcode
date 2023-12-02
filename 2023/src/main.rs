mod day01;
mod day02;
mod puzzle;
mod runner;

fn main() {
    let day = std::env::args().nth(1).unwrap_or_else(|| "".to_string());

    if day.is_empty() {
        eprintln!("Please provide a day to run e.g. `cargo run 1`");
        std::process::exit(1);
    }

    runner::run(&day);
}
