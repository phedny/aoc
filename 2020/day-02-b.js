const fs = require('fs');

const result = fs.readFileSync('./inputs/2/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/(\d+)-(\d+) (\w): (\w+)/))
	.map(([_, pos1, pos2, letter, password]) => [password.split('')[Number.parseInt(pos1) - 1], password.split('')[Number.parseInt(pos2) - 1], letter])
	.filter(([pos1, pos2, letter]) => (pos1 === letter || pos2 === letter) && pos1 !== pos2)
	.length;

console.log(result);
