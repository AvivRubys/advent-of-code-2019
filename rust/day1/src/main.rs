use std::fs::File;
use std::io::{self, BufRead};

fn fuel_by_mass(mass: i32) -> i32 {
    mass / 3 - 2
}

fn fuel_by_mass_recursive(mass: i32) -> i32 {
    let fuel = fuel_by_mass(mass);
    // println!("Mass - {}, Fuel - {}", mass, fuel);

    if fuel <= 0 {
        0
    } else {
        fuel_by_mass_recursive(fuel) + fuel
    }
}

fn to_int(string: String) -> i32 {
    let mass: i32 = string.parse().unwrap();
    mass
}

fn main() {
    match File::open("./input") {
        Ok(file) => {
            let lines = io::BufReader::new(file).lines();
            let masses: Vec<i32> = lines.map(|line| line.unwrap()).map(to_int).collect();
            let sum = masses.iter().fold(0, |acc, &mass| acc + fuel_by_mass(mass));
            let sum_including_fuel = masses
                .iter()
                .fold(0, |acc, &mass| acc + fuel_by_mass_recursive(mass));
            println!("Sum is {}", sum);
            println!("Sum including fuel is {}", sum_including_fuel);
        }
        Err(error) => eprintln!("Couldn't open file - {}", error),
    }
}
