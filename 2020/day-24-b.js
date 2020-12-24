const fs = require('fs');

const directions = { e: [0, 1], se: [-1, 0], sw: [-1, -1], w: [0, -1], nw: [1, 0], ne: [1, 1] };

const tiles = fs.readFileSync('./inputs/24/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => [...line.matchAll(/[ns]?[ew]/g)]
		.map(([m]) => directions[m])
		.reduce(([a, b], [dA, dB]) => [a + dA, b + dB], [0, 0])
	)
	.map(coord => coord.join(','));

let blackTiles = [...new Set(tiles)].filter(tile => tiles.filter(t2 => tile === t2).length % 2 === 1);

function day(black) {
	const adjacents = black.map(coord => coord.split(',').map(n => Number.parseInt(n)))
		.flatMap(([a, b]) => Object.values(directions).map(([dA, dB]) => [a + dA, b + dB]))
		.map(coord => coord.join(','));
	const toBlack = [...new Set(adjacents)].filter(coord => {
		const count = adjacents.filter(a => a === coord).length;
		return !black.includes(coord) && count === 2;
	});
	const toWhite = black.filter(coord => {
		const count = adjacents.filter(a => a === coord).length;
		return count === 0 || count > 2;
	});

	return [...black.filter(coord => !toWhite.includes(coord)), ...toBlack];
}

for (let i = 0; i < 100; i++) {
	blackTiles = day(blackTiles);
}

console.log(blackTiles.length);
