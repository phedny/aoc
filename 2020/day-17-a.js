const fs = require('fs');

let coords = fs.readFileSync('./inputs/17/real.txt', { encoding: 'utf8' })
	.split('\n')
	.flatMap((line, i) => line.split('').map((c, j) => [c, i, j]))
	.filter(([c]) => c === '#')
	.map(([_, i, j]) => [0, i, j]);

function cycle(coords) {
	const prevCoords = coords.map(coord => coord.join(','));
	const allCounts = coords.flatMap(c => c.map(a => [a - 1, a, a + 1])
		.reduce((acc, list) => acc.flatMap(prefix => list.map(suffix => [...prefix, suffix])), [[]]))
		.map(coord => coord.join(','));
	return [...new Set(allCounts)].map(coord => [coord, allCounts.filter(cc => coord === cc).length])
		.filter(([coord, count]) => prevCoords.includes(coord) ? count === 3 || count === 4 : count === 3)
		.map(([coord]) => coord.split(',').map(c => Number.parseInt(c)));
}

for (let i = 0; i < 6; i++) coords = cycle(coords);

console.log(coords.length);
