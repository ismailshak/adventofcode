mod day01;

fn main() {
    let day = std::env::args()
        .nth(1)
        .expect("Pick a specific day to execute: cargo run 01");

    match day.as_str() {
        "1" | "01" => day01::execute(),
        x => println!("{} is not implemented yet", x),
    }
}
