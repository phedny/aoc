const fs = require('fs');

const numbers = [
	0,
	...(
		fs.readFileSync('./inputs/10/real.txt', { encoding: 'utf8' })
			.split('\n')
			.map(line => Number.parseInt(line))
			.sort((a, b) => a - b)
	)
];

const differences = numbers.map((n, i) => numbers[i + 1] - n || 3);

const result = differences.filter(d => d === 1).length * differences.filter(d => d === 3).length;

console.log(result);
