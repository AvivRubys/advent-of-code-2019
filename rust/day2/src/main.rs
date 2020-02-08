use std::fs::read_to_string;

fn evaluate(instructions: &mut Vec<i64>) {
    let mut i = 0;
    loop {
        match instructions[i] {
            1 => {
                let op1 = instructions[i + 1] as usize;
                let op2 = instructions[i + 2] as usize;
                let op3 = instructions[i + 3] as usize;

                instructions[op3] = instructions[op1] + instructions[op2];
                i += 4;
            }
            2 => {
                let op1 = instructions[i + 1] as usize;
                let op2 = instructions[i + 2] as usize;
                let op3 = instructions[i + 3] as usize;

                instructions[op3] = instructions[op1] * instructions[op2];
                i += 4;
            }
            99 => break,
            _ => {
                println!("This should not happen");
                break;
            }
        }
    }
}

fn main() {
    match read_to_string("./input") {
        Ok(contents) => {
            let instructions: Vec<i64> = contents
                .split(',')
                .map(|op| str::parse::<i64>(op).unwrap())
                .collect();

            'attempts: for noun in 0..99 {
                for verb in 0..99 {
                    let mut instructions_clone = instructions.clone();
                    instructions_clone[1] = noun;
                    instructions_clone[2] = verb;
                    evaluate(&mut instructions_clone);

                    if instructions_clone[0] == 19690720 {
                        println!("Result is 100*{}+{}={}", noun, verb, 100 * noun + verb);
                        break 'attempts;
                    }
                }
            }
        }
        Err(error) => eprintln!("Couldn't read file - {}", error),
    }
}
