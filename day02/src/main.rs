use std::fs;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    let mut part_one = 0;
    let mut part_two = 0;
    for line in contents.lines() {
        let first = line.chars().nth(0).unwrap();
        let last = line.chars().nth(2).unwrap();

        // Part One
        if last == 'X' {
            part_one += 1;
        } else if last == 'Y' {
            part_one += 2;
        } else {
            part_one += 3;
        }
        if (first == 'A' && last == 'Y')
            || (first == 'B' && last == 'Z')
            || (first == 'C' && last == 'X')
        {
            // Won
            part_one += 6;
        } else if (first == 'A' && last == 'X')
            || (first == 'B' && last == 'Y')
            || (first == 'C' && last == 'Z')
        {
            // Tied
            part_one += 3;
        }

        // Part Two
        if last == 'X' {
            // Lost
            if first == 'A' {
                part_two += 3;
            } else if first == 'B' {
                part_two += 1;
            } else {
                part_two += 2;
            }
        } else if last == 'Y' {
            // Tied
            part_two += 3;
            if first == 'A' {
                part_two += 1;
            } else if first == 'B' {
                part_two += 2;
            } else {
                part_two += 3;
            }
        } else {
            // Won
            part_two += 6;
            if first == 'A' {
                part_two += 2;
            } else if first == 'B' {
                part_two += 3;
            } else {
                part_two += 1;
            }
        }
    }
    println!("Part One: {}", part_one);
    println!("Part Two: {}", part_two);
}
