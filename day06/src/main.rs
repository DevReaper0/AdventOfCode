use std::collections::HashSet;
use std::fs;

fn part_one(buffer: &str) -> i32 {
    let mut last_four = String::new();
    for (i, c) in buffer.chars().enumerate() {
        last_four.push(c);
        if last_four.len() > 4 {
            last_four.remove(0);
        }
        if last_four.chars().collect::<HashSet<char>>().len() == 4 {
            return i as i32 + 1;
        }
    }
    return -1;
}

fn part_two(buffer: &str) -> i32 {
    let mut last_fourteen = String::new();
    for (i, c) in buffer.chars().enumerate() {
        last_fourteen.push(c);
        if last_fourteen.len() > 14 {
            last_fourteen.remove(0);
        }
        if last_fourteen.chars().collect::<HashSet<char>>().len() == 14 {
            return i as i32 + 1;
        }
    }
    return -1;
}

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    println!("Part One: {}", part_one(&contents));
    println!("Part Two: {}", part_two(&contents));
}
