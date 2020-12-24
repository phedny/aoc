const fs = require('fs');

const directions = { e: [0, 1], se: [-1, 0], sw: [-1, -1], w: [0, -1], nw: [1, 0], ne: [1, 1] };

const tiles = fs.readFileSync('./inputs/24/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => [...line.matchAll(/[ns]?[ew]/g)]
		.map(([m]) => directions[m])
		.reduce(([a, b], [dA, dB]) => [a + dA, b + dB], [0, 0])
	)
	.map(coord => coord.join(','));

const result = [...new Set(tiles)].filter(tile => tiles.filter(t2 => tile === t2).length % 2 === 1).length;

console.log(result);
