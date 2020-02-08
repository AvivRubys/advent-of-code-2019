extern crate simple_error;

use simple_error::SimpleError;
use std::cmp::Ordering;
use std::fs::read_to_string;
use std::str::FromStr;

#[derive(Debug, Clone, Copy)]
enum Movement {
    Up(u32),
    Down(u32),
    Left(u32),
    Right(u32),
}

impl FromStr for Movement {
    type Err = SimpleError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let direction = s.chars().nth(0);
        let distance = s[1..].parse::<u32>().unwrap();
        use Movement::*;
        match direction {
            Some('U') => Ok(Up(distance)),
            Some('D') => Ok(Down(distance)),
            Some('L') => Ok(Left(distance)),
            Some('R') => Ok(Right(distance)),
            Some(ch) => Err(SimpleError::new(format!("Unexpected char {}", ch))),
            None => Err(SimpleError::new("Parsing error")),
        }
    }
}

#[derive(Debug, Clone, Copy)]
struct Point {
    x: i64,
    y: i64,
}

impl Point {
    pub fn add(self, movement: Movement) -> Self {
        match movement {
            Movement::Up(n) => Point {
                y: self.y + (n as i64),
                ..self
            },
            Movement::Down(n) => Point {
                y: self.y - (n as i64),
                ..self
            },
            Movement::Right(n) => Point {
                x: self.x + (n as i64),
                ..self
            },
            Movement::Left(n) => Point {
                x: self.x - (n as i64),
                ..self
            },
        }
    }

    pub fn distance(self, other: Self) -> i64 {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

#[derive(Debug, Clone, Copy)]
struct Line(Point, Point);

impl Line {
    pub fn intersects(self, other: Self) -> Option<Point> {
        let a1 = self.1.y - self.0.y;
        let b1 = self.0.x - self.1.x;
        let c1 = a1 * self.0.x + b1 * self.0.y;
        let a2 = other.1.y - other.0.y;
        let b2 = other.0.x - other.1.x;
        let c2 = a2 * other.0.x + b2 * other.0.y;
        let delta = a1 * b2 - a2 * b1;
        if delta == 0 {
            return None;
        }

        let x = (b2 * c1 - b1 * c2) / delta;
        let y = (a1 * c2 - a2 * c1) / delta;

        if !(is_between(x, self.0.x, self.1.x)
            && is_between(y, self.0.y, self.1.y)
            && is_between(x, other.0.x, other.1.x)
            && is_between(y, other.0.y, other.1.y))
        {
            None
        } else {
            Some(Point { x: x, y: y })
        }
    }
}

fn is_between(n: i64, b1: i64, b2: i64) -> bool {
    b1 <= n && n <= b2 || b2 <= n && n <= b1
}

fn wire_to_movements(wire: &str) -> Vec<Movement> {
    wire.split(',')
        .map(|x| str::parse::<Movement>(x).unwrap())
        .collect()
}

fn movements_to_lines<'a>(initial_point: Point, moves: &'a [Movement]) -> Vec<Line> {
    let mut lines = Vec::with_capacity(moves.len());
    let mut last_point = initial_point;

    for &m in moves {
        let next_point = last_point.add(m);
        lines.push(Line(last_point, next_point));
        last_point = next_point;
    }

    lines
}

fn find_intersections(wire1: Vec<Line>, wire2: Vec<Line>) -> Vec<Point> {
    let mut intersections = Vec::new();
    for line1 in wire1 {
        for &line2 in &wire2 {
            if let Some(intersection_point) = line1.intersects(line2) {
                intersections.push(intersection_point);
            }
        }
    }

    intersections
}

const POINT_ZERO: Point = Point { x: 0, y: 0 };

fn main() {
    if let Ok(file) = read_to_string("./input") {
        let movements: Vec<Vec<Movement>> = file.split('\n').map(wire_to_movements).collect();
        let wire1 = movements_to_lines(POINT_ZERO, &movements[0]);
        let wire2 = movements_to_lines(POINT_ZERO, &movements[1]);
        let intersections = find_intersections(wire1, wire2);
        let min_intersection = intersections
            .iter()
            .filter(|p| p.x != 0 || p.y != 0)
            .min_by(|p1, p2| {
                let d1 = p1.distance(POINT_ZERO);
                let d2 = p2.distance(POINT_ZERO);

                if d2 < d2 {
                    Ordering::Less
                } else if d1 > d2 {
                    Ordering::Greater
                } else {
                    Ordering::Equal
                }
            });

        println!("Found {:?}", min_intersection);
        println!(
            "Distance is {}",
            min_intersection.unwrap().distance(POINT_ZERO)
        );
    } else {
        eprintln!("Error opening file")
    }
}
