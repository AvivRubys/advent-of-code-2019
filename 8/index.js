const fs = require('fs');
const R = require('ramda');

const WIDTH = 25,
    HEIGHT = 6;

const input = fs
    .readFileSync('./input.txt')
    .toString()
    .split('');
const result = R.pipe(
    R.splitEvery(WIDTH * HEIGHT),
    R.map(R.countBy(x => x)),
    R.reduce(
        R.minBy(layer => layer['0']),
        {'0': Infinity}
    ),
    val => console.log(val['1'] * val['2'])
)(input);
