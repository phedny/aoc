const fs = require('fs');

const tiles = fs.readFileSync('./inputs/20/real.txt', { encoding: 'utf8' })
	.split('\n\n')
	.map(data => {
		const lines = data.split('\n');
		const [_, firstLine] = lines.shift().match(/Tile (\d+):/);
		const cells = lines.map(line => line.split(''));

		return {
			tileNumber: Number.parseInt(firstLine),
			cells,
			borders: [
				cells[0],
				cells.map(row => row[row.length - 1]),
				[...cells[cells.length - 1]].reverse(),
				[...cells.map(row => row[0])].reverse()
			]
		};
	});

function rotate(tile, count) {
	if (count === 0) return tile;
	const cells = tile.cells.map((_i, i) => tile.cells.map((_j, j) => tile.cells[j][tile.cells.length - i - 1]));
	return rotate({
		tileNumber: tile.tileNumber,
		cells,
		borders: [
			cells[0],
			cells.map(row => row[row.length - 1]),
			[...cells[cells.length - 1]].reverse(),
			[...cells.map(row => row[0])].reverse()
		]
	}, count - 1);
}

function flip(tile) {
	const cells = [...tile.cells].reverse();
	return {
		tileNumber: tile.tileNumber,
		cells,
		borders: [
			cells[0],
			cells.map(row => row[row.length - 1]),
			[...cells[cells.length - 1]].reverse(),
			[...cells.map(row => row[0])].reverse()
		]
	};
}

function borderToNum(border) {
	return Number.parseInt(border.map(c => c === '#' ? 1 : 0).join(''), 2);
}

function tileToNum(tile) {
	return {
		tileNumber: tile.tileNumber,
		cells: tile.cells,
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

const theMap = [];
for (let i = 0; i < size; i++) {
	const blocks = solution.slice(i * size, (i + 1) * size)
		.map(tile => tile.cells.slice(1, tile.cells.length - 1).map(row => row.slice(1, row.length - 1).join('')))
	theMap.splice(theMap.length, 0, ...blocks[0].map((_, j) => blocks.map(b => b[j]).join('').split('')));
}

const allMaps = [{ cells: theMap }].flatMap(tile => [...Array(4).keys()].map(r => rotate(tile, r)))
	.flatMap(tile => [tile, flip(tile)])
	.map(tile => tile.cells);

let seaMonster = [
	'                  # '.split('').flatMap((c, i) => c === '#' ? [[0, i]] : []),
	'#    ##    ##    ###'.split('').flatMap((c, i) => c === '#' ? [[1, i]] : []),
	' #  #  #  #  #  #   '.split('').flatMap((c, i) => c === '#' ? [[2, i]] : [])
].flat();

let result;
allMaps.forEach(thisMap => {
	let hasOne = false;
	for (let i = 0; i < thisMap.length - 3; i++) {
		for (let j = 0; j < thisMap.length - 20; j++) {
			if (seaMonster.every(([x, y]) => thisMap[i + x][j + y] === '#')) {
				hasOne = true;

				seaMonster.forEach(([x, y]) => thisMap[i + x][j + y] = 'O');
			}
		}
	}

	if (hasOne) {
		result = thisMap.flatMap(line => line.filter(c => c === '#').length).reduce((a, b) => a + b, 0);
	}
});

console.log(result);
