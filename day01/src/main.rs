use std::fs;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    let mut elves: Vec<i32> = Vec::new();
    for line in contents.lines() {
        if line.len() > 0 {
            if elves.len() == 0 {
                elves.push(0);
            }
            let i = elves.len() - 1;
            if let Some(elem) = elves.get_mut(i) {
                *elem += line.parse::<i32>().unwrap();
            }
        } else {
            elves.push(0);
        }
    }

    elves.sort_by(|a, b| b.cmp(a));

    println!("Part One: {}", elves[0]);
    println!("Part Two: {}", (elves[0] + elves[1] + elves[2]));
}
