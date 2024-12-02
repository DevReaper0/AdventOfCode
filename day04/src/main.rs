use std::fs;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    let mut part_one = Vec::new();
    let mut part_two = Vec::new();
    for line in contents.lines() {
        let mut ranges = line.split(",");
        let first = ranges.next().unwrap();
        let second = ranges.next().unwrap();

        let mut first_split = first.split("-");
        let first_min = first_split.next().unwrap().parse::<i32>().unwrap();
        let first_max = first_split.next().unwrap().parse::<i32>().unwrap();

        let mut second_split = second.split("-");
        let second_min = second_split.next().unwrap().parse::<i32>().unwrap();
        let second_max = second_split.next().unwrap().parse::<i32>().unwrap();

        if (first_min <= second_min && first_max >= second_max)
            || (second_min <= first_min && second_max >= first_max)
        {
            part_one.push((first, second));
        }
        if first_min <= second_max && second_min <= first_max {
            part_two.push((first, second));
        }
    }
    println!("Part One: {}", part_one.len());
    println!("Part Two: {}", part_two.len());
}
