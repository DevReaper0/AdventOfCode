use std::fs;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    let mut crates: Vec<Vec<char>> = Vec::new();
    for line in contents.lines() {
        if !line.contains("[") {
            crates = vec![
                vec![];
                line.split_whitespace()
                    .last()
                    .unwrap()
                    .parse::<i32>()
                    .unwrap() as usize
            ];
            break;
        }
    }
    for line in contents.lines() {
        if line.contains("[") {
            let mut i = 0;
            let mut mode = 0;
            for c in (line.to_owned() + " ").chars() {
                mode += 1;
                if mode == 2 {
                    if c != ' ' {
                        crates[i].insert(0, c);
                    }
                    i += 1;
                }
                if mode == 4 {
                    mode = 0;
                }
            }
        }
    }
    let mut part_one = crates.clone();
    let mut part_two = crates.clone();
    for line in contents.lines() {
        if line.contains("move ") {
            let mut amount = 0;
            let mut from = 0;
            let mut to = 0;
            let mut mode = 1;
            for part in line.split_whitespace() {
                if part == "move" {
                    mode = 1;
                    continue;
                } else if part == "from" {
                    mode = 2;
                    continue;
                } else if part == "to" {
                    mode = 3;
                    continue;
                }

                if mode == 1 {
                    amount = part.parse::<i32>().unwrap();
                } else if mode == 2 {
                    from = part.parse::<i32>().unwrap();
                } else if mode == 3 {
                    to = part.parse::<i32>().unwrap();
                }
            }
            for _ in 0..amount {
                let cr = part_one[(from - 1) as usize].pop().unwrap();
                part_one[(to - 1) as usize].push(cr);
            }
            let index = part_two[(to - 1) as usize].len();
            for _ in 0..amount {
                let cr = part_two[(from - 1) as usize].pop().unwrap();
                part_two[(to - 1) as usize].insert(index, cr);
            }
        }
    }
    let mut part_one_result = String::new();
    for crate_stack in part_one {
        if crate_stack.len() > 0 {
            part_one_result.push(crate_stack[crate_stack.len() - 1]);
        }
    }
    let mut part_two_result = String::new();
    for crate_stack in part_two {
        if crate_stack.len() > 0 {
            part_two_result.push(crate_stack[crate_stack.len() - 1]);
        }
    }
    println!("Part One: {}", part_one_result);
    println!("Part Two: {}", part_two_result);
}
