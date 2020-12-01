const fs = require('fs');

const numbers = fs.readFileSync('./inputs/1/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => Number.parseInt(line));

const allPairs = numbers.flatMap((n1, i) => numbers.slice(i).map(n2 => [n1, n2]));

const result = allPairs.filter(([n1, n2]) => n1 + n2 === 2020)
	.map(([n1, n2]) => n1 * n2);

console.log(result);
