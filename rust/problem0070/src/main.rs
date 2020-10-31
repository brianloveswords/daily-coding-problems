use std::env;

fn perfect_number(n: usize, command: Command) -> usize {
    if n <= 0 || n > 59271 {
        panic!("valid range for n is [1, 59271] got {}", n)
    }
    let mut i = 0;
    let mut current = 19;
    let mut next_power = 100;

    let handle = match command {
        Command::Graph => |i: usize, v: usize| println!("{},{}", i, v),
        Command::None => |_i: usize, _v: usize| (),
    };

    loop {
        if is_perfect(current) {
            handle(i, current);
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

#[derive(Debug)]
enum Command {
    Graph,
    None,
}

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();

    eprintln!("args {:?}", args);

    let n = match args[0].parse::<usize>() {
        Ok(n) => n,
        Err(err) => panic!(err),
    };

    let command = match args
        .get(1)
        .map(|x| x.to_owned())
        .map(|x| x.to_ascii_lowercase())
        .unwrap_or("none".to_owned())
        .as_ref()
    {
        "graph" => Command::Graph,
        "none" => Command::None,
        _ => panic!("expected `graph` or `none`"),
    };

    eprintln!("command {:?}", command);

    eprintln!("{}", perfect_number(n, command));
}
