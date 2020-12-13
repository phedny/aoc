const fs = require('fs');

const [earliestBus, busList] = fs.readFileSync('./inputs/13/real.txt', { encoding: 'utf8' }).split('\n');
const [minutes, line] = busList.split(',')
	.filter(bus => bus !== 'x')
	.map(bus => Number.parseInt(bus))
	.map(bus => [bus, bus - (earliestBus % bus)])
	.sort(([_a, a], [_b, b]) => a - b)[0];

console.log(minutes * line);
