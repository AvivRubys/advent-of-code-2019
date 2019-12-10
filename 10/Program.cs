using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;

// ReSharper disable once CheckNamespace
namespace Ten
{
    static class Utils
    {
        public static int Gcd(int a, int b)
        {
            if (b == 0) return a;
            return Gcd(b, a % b);
        }
    }
    
    struct Point
    {
        public int X { get; set; }
        public int Y { get; set; }

        public Point(int x, int y)
        {
            X = x;
            Y = y;
        }
        
        public Point Reduce()
        {
            var gcd = Math.Abs(Utils.Gcd(X, Y));

            if (gcd == 1) return this;

            var newPoint = new Point(X / gcd, Y / gcd);
            
            if (X < 0)
            {
                newPoint.X = -Math.Abs(newPoint.X);
            }
            
            if (Y < 0)
            {
                newPoint.Y = -Math.Abs(newPoint.Y);
            }

            return newPoint;
        }

        public double AngleTo(Point p)
        {
            return Math.Atan2(p.X - X, p.Y - Y) * 180 / Math.PI;
        }

        public int DistanceFrom(Point p)
        {
            var diff = this - p;
            return Math.Abs(diff.X) + Math.Abs(diff.Y);
        }

        public override string ToString()
        {
            return $"({X},{Y})";
        }
        
        public static bool operator ==(Point a, Point b)
        {
            return a.Equals(b);
        }
        
        public static bool operator !=(Point a, Point b)
        {
            return !a.Equals(b);
        }
        
        public static Point operator *(Point a, int factor)
        {
            return new Point
            {
                X = a.X * factor,
                Y = a.Y * factor
            };
        }
        
        public static Point operator +(Point a, Point b)
        {
            return new Point
            {
                X = a.X + b.X,
                Y = a.Y + b.Y
            };
        }
        
        public static Point operator -(Point a, Point b)
        {
            return new Point
            {
                X = a.X - b.X,
                Y = a.Y - b.Y
            };
        }

        
        public bool Equals(Point other)
        {
            return X == other.X && Y == other.Y;
        }

        public override bool Equals(object obj)
        {
            return obj is Point other && Equals(other);
        }

        public override int GetHashCode()
        {
            return ToString().GetHashCode();
        }
    }

    class Program
    {
        static int VisibleFrom(Point asteroid, ICollection<Point> asteroids, int height, int width)
        {
            var isHidden = new bool[height,width];

            foreach (var otherAstroid in asteroids)
            {
                if (asteroid == otherAstroid) continue;
                var distance = (otherAstroid - asteroid).Reduce();

                var factor = 1;
                while (true)
                {
                    var unseenPoint = otherAstroid + distance * factor;
                    if (unseenPoint.X >= width || unseenPoint.X < 0 ||
                        unseenPoint.Y >= height || unseenPoint.Y < 0) break;
                    
                    isHidden[unseenPoint.X, unseenPoint.Y] = true;
                    factor++;
                }
            }

            return asteroids.Count(a => !isHidden[a.X, a.Y]) - 1;
        }

        static void Main()
        {
            var input = File.ReadAllLines("./input.txt");
            var height = input.Length;
            var width = input[0].Length;
            var asteroids = input.SelectMany((row, y) => row.Select((val, x) => new {val, point = new Point(x, y)}))
                .Where(p => p.val == '#').Select(p => p.point).ToHashSet();

            var maxAsteroid = asteroids.Select(asteroid => new {
                Asteroid = asteroid,
                Count = VisibleFrom(asteroid, asteroids, height, width)
            }).Aggregate((max, val) => val.Count > max.Count ? val : max);
            
            Console.WriteLine(maxAsteroid);
            
            // Part 2
            var rootPoint = maxAsteroid.Asteroid;
            var sortedAngles = asteroids.Select(asteroid => (asteroid, rootPoint.AngleTo(asteroid)))
                .Where(tup => asteroids.Contains(tup.Item1))
                .GroupBy(tup => tup.Item2, tup => tup.Item1)
                .OrderByDescending(grp => grp.Key)
                .Select(grp => new Queue<Point>(grp.OrderBy(p => rootPoint.DistanceFrom(p))))
                .ToArray();

            var foundPoints = false;
            var i = 1;
            do {
                foundPoints = false;
                foreach (var angle in sortedAngles)
                {
                    if (angle.TryDequeue(out var point))
                    {
                        if (i == 200)
                        {
                            Console.WriteLine("200th asteroid is {0}, result = {1}", point, point.X * 100 + point.Y);
                        }
                        foundPoints = true;
                        i++;
                    }
                }
            }
            while (foundPoints) ;
        }
    }
}