const fs = require('fs');

const result = fs.readFileSync('./inputs/2/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/(\d+)-(\d+) (\w): (\w+)/))
	.map(([_, min, max, letter, password]) => [Number.parseInt(min), Number.parseInt(max), password.split('').filter(char => char === letter).length])
	.filter(([min, max, count]) => min <= count && max >= count)
	.length;

console.log(result);
