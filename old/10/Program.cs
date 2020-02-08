using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

// ReSharper disable once CheckNamespace
namespace Ten
{
    struct Point
    {
        public int X { get; set; }
        public int Y { get; set; }

        public Point(int x, int y)
        {
            X = x;
            Y = y;
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
        
        public override string ToString()
        {
            return $"({X},{Y})";
        }
    }

    class Program
    {
        static void Main()
        {
            var input = File.ReadAllLines("./input.txt");
            var asteroids = input.SelectMany((row, y) => row.Select((val, x) => new {val, point = new Point(x, y)}))
                .Where(p => p.val == '#')
                .Select(p => p.point)
                .ToArray();

            var maxAsteroid = asteroids.Select(asteroid => new
            {
                Asteroid = asteroid,
                Count = asteroids.Select(asteroid.AngleTo).Distinct().Count()
            }).Aggregate((p, c) => c.Count > p.Count ? c : p);

            Console.WriteLine(maxAsteroid);

            // Part 2
            var rootPoint = maxAsteroid.Asteroid;
            var sortedAngles = asteroids.Select(asteroid => (asteroid, rootPoint.AngleTo(asteroid)))
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