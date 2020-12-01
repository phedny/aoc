const fs = require('fs');

const numbers = fs.readFileSync('./inputs/1/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => Number.parseInt(line));

const allTriples = [numbers, numbers, numbers].reduce((acc, list) => acc.flatMap(prefix => list.map(suffix => [...prefix, suffix])), [[]]);

const result = allTriples.filter(nn => nn.reduce((acc, n) => acc + n, 0) === 2020)
	.map(nn => nn.reduce((acc, n) => acc * n, 1));

console.log(result[0]);
