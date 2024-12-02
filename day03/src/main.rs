use std::fs;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    let mut part_one = 0;
    for line in contents.lines() {
        let first = line.chars().take(line.len() / 2).collect::<Vec<_>>();
        let second = line.chars().skip(line.len() / 2).collect::<Vec<_>>();

        for item in first {
            if second.contains(&item) {
                part_one += match item as u8 {
                    b'a'..=b'z' => ((item as u8) - b'a') as u32 + 1,
                    b'A'..=b'Z' => ((item as u8) - b'A') as u32 + 27,
                    _ => panic!("Unexpected character"),
                };
                break;
            }
        }
    }

    let mut part_two = 0;
    let mut lines = contents.lines();
    while let Some(first) = lines.next() {
        let second = lines.next().unwrap();
        let third = lines.next().unwrap();

        for item in first.chars() {
            if second.contains(item) && third.contains(item) {
                part_two += match item as u8 {
                    b'a'..=b'z' => ((item as u8) - b'a') as u32 + 1,
                    b'A'..=b'Z' => ((item as u8) - b'A') as u32 + 27,
                    _ => panic!("Unexpected character"),
                };
                break;
            }
        }
    }

    println!("Part One: {}", part_one);
    println!("Part Two: {}", part_two);
}
