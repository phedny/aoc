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

const pathsToNumber = { 0: 1 };

numbers.forEach(curr => {
	[1, 2, 3].map(a => curr + a)
		.filter(next => numbers.includes(next))
		.forEach(next => pathsToNumber[next] = (pathsToNumber[next] || 0) + pathsToNumber[curr])
});

console.log(pathsToNumber[numbers[numbers.length - 1]]);
