const fs = require('fs');

const tiles = fs.readFileSync('./inputs/20/real.txt', { encoding: 'utf8' })
	.split('\n\n')
	.map(data => {
		const lines = data.split('\n');
		const [_, firstLine] = lines.shift().match(/Tile (\d+):/);
		const cells = lines.map(line => line.split(''));

		return {
			tileNumber: Number.parseInt(firstLine),
			borders: [
				cells[0],
				cells.map(row => row[row.length - 1]),
				[...cells[cells.length - 1]].reverse(),
				[...cells.map(row => row[0])].reverse()
			]
		};
	});

function rotate(tile, count) {
	return {
		tileNumber: tile.tileNumber,
		borders: [ ...tile.borders.slice(count, 4), ...tile.borders.slice(0, count) ]
	};
}

function flip(tile) {
	return {
		tileNumber: tile.tileNumber,
		borders: [tile.borders[0], tile.borders[3], tile.borders[2], tile.borders[1]].map(border => [...border].reverse())
	};
}

function borderToNum(border) {
	return Number.parseInt(border.map(c => c === '#' ? 1 : 0).join(''), 2);
}

function tileToNum(tile) {
	return {
		tileNumber: tile.tileNumber,
		borders: tile.borders.map(border => borderToNum(border))
	};
}

const tileOptions = tiles.flatMap(tile => [...Array(4).keys()].map(r => rotate(tile, r)))
	.flatMap(tile => [tile, flip(tile)]);

const size = Math.sqrt(tiles.length);
function find(solution) {
	if (solution.length === tiles.length) return solution;
	for (const tile of tileOptions) {
		if (solution.some(solTile => solTile.tileNumber === tile.tileNumber)) continue;
		if (solution.length >= size && borderToNum(solution[solution.length - size].borders[2]) !== borderToNum([...tile.borders[0]].reverse())) continue;
		if (solution.length % size !== 0 && borderToNum(solution[solution.length - 1].borders[1]) !== borderToNum([...tile.borders[3]].reverse())) continue;
		const found = find([...solution, tile]);
		if (found) return found;
	}
}

const solution = find([]);
const result = solution[0].tileNumber * solution[size - 1].tileNumber * solution[solution.length - size].tileNumber * solution[solution.length - 1].tileNumber;

console.log(result);
