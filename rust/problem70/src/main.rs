fn perfect_number(n: usize) -> usize {
    if n <= 0 || n > 59271 {
        panic!("valid range for n is [1, 59271] got {}", n)
    }
    let mut i = 0;
    let mut current = 19;
    let mut next_power = 100;

    loop {
        if is_perfect(current) {
            i += 1
        }
        if i == n {
            return current;
        }
        if current == next_power {
            eprintln!("debug: passed {}", next_power);
            next_power = next_power * 10;
        }
        current += 1
    }
}

fn is_perfect(n: usize) -> bool {
    let mut current = n;
    let mut sum = 0;
    loop {
        let remaining_places = current / 10;
        if remaining_places == 0 {
            sum += current;
            return sum == 10;
        }
        sum += current - remaining_places * 10;
        current = remaining_places;
    }
}

fn main() {
    eprintln!("{}", perfect_number(59271));
}
