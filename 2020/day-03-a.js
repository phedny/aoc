const fs = require('fs');

const result = fs.readFileSync('./inputs/3/real.txt', { encoding: 'utf8' })
	.split('\n')
	.filter((line, i) => line.split('')[(3 * i) % line.length] === '#')
	.length;

console.log(result);
