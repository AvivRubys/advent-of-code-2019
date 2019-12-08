const fs = require('fs');
const os = require('os');
const R = require('ramda');

const WIDTH = 25,
    HEIGHT = 6;

const layers = R.pipe(R.split(''), R.splitEvery(WIDTH * HEIGHT))(fs.readFileSync('./input.txt').toString());

// part 1
R.pipe(
    R.map(R.countBy(x => x)),
    R.reduce(
        R.minBy(layer => layer['0']),
        {'0': Infinity}
    ),
    val => console.log(val['1'] * val['2'])
)(layers);

// part 2
const zipMany = list => R.reduce(R.pipe(R.zip, R.map(R.flatten)), R.head(list), R.tail(list));
R.pipe(
    zipMany,
    R.map(R.find(a => a !== '2')),
    R.map(x => (x === '0' ? ' ' : '#')),
    R.splitEvery(WIDTH),
    R.forEach(R.pipe(R.join(' '), console.log))
)(layers);
