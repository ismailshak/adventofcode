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

fn part2() -> u32 {
    include_str!("./input.txt")
        .lines()
        .map(|l| {
            let parsed = l
                .to_string()
                // Wrapping with string equiv. so that overlaps like `sevenine` are handled
                .replace("one", "one1one")
                .replace("two", "two2two")
                .replace("three", "three3three")
                .replace("four", "four4four")
                .replace("five", "five5five")
                .replace("six", "six6six")
                .replace("seven", "seven7seven")
                .replace("eight", "eight8eight")
                .replace("nine", "nine9nine")
                .chars()
                .filter_map(|x| x.to_digit(10))
                .collect::<Vec<u32>>();

            let first = parsed.first().expect("Failed to find first");
            let last = parsed.last().expect("Failed to find last");

            first * 10 + last
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
        "Calibration Sum pt.1: {} ({:.3}ms)", // 3 decimal places printed
        calibration_sum,
        duration as f64 / 1_000_000_f64
    );

    let start = std::time::Instant::now();
    let calibration_sum = part2();
    let duration = start.elapsed().as_nanos();

    println!(
        "Calibration Sum pt.2: {} ({:.3}ms)",
        calibration_sum,
        duration as f64 / 1_000_000_f64
    );
}
