const fs = require('fs');

function count(right, down) {
	return fs.readFileSync('./inputs/3/real.txt', { encoding: 'utf8' })
		.split('\n')
		.filter((line, i) => i % down === 0)
		.filter((line, i) => line.split('')[(right * i) % line.length] === '#')
		.length;
}

console.log(count(1, 1) * count(3, 1) * count(5, 1) * count(7, 1) * count(1, 2));
