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

function flipHorizontal(tile) {
	return {
		tileNumber: tile.tileNumber,
		borders: [tile.borders[0], tile.borders[3], tile.borders[2], tile.borders[1]].map(border => [...border].reverse())
	};
}

function flipVertical(tile) {
	return {
		tileNumber: tile.tileNumber,
		borders: [tile.borders[2], tile.borders[1], tile.borders[0], tile.borders[3]].map(border => [...border].reverse())
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
	.flatMap(tile => [tile, flipHorizontal(tile), flipVertical(tile)]);

const tilesWithMatches = tiles.map(tile => {
	const borderMatches = tile.borders.map(border => borderToNum([...border].reverse())).map(border =>
		tileOptions.filter(tile2 => tile.tileNumber !== tile2.tileNumber)
			.filter(tile2 => tile2.borders.map(border => borderToNum(border)).includes(border))
	);
	return { borderMatches, ...tile };
});

const result = tilesWithMatches.filter(tile => tile.borderMatches.filter(m => m.length === 0).length === 2)
	.map(tile => tile.tileNumber)
	.reduce((a, b) => a * b, 1);

console.log(result);
