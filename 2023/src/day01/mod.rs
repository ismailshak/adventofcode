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

pub fn execute() {
    println!("Day 1: Trebuchet?!");
    println!("------------------");

    let start = std::time::Instant::now();
    let calibration_sum = part1();
    let duration = start.elapsed().as_nanos();

    println!(
        "Calibration Sum: {} ({:.3}ms)", // 3 decimal places printed
        calibration_sum,
        duration as f64 / 1_000_000_f64
    );
}
