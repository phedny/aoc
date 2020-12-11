const fs = require('fs');

const map = fs.readFileSync('./inputs/11/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => ['.', ...line.split(''), '.']);

map.unshift('.'.repeat(map[0].length).split(''));
map.push(map[0]);

function* gen(input) {
	let map = input;
	while (true) {
		let hasChange = false;
		const nextMap = new Array(map.length);
		nextMap[0] = nextMap[nextMap.length - 1] = map[0];
		for (let i = 1; i < map.length - 1; i++) {
			nextMap[i] = [...map[0]];
			for (let j = 1; j < map[0].length - 1; j++) {
				 const adjacent = [
				 	map[i - 1][j - 1],
				 	map[i - 1][j],
				 	map[i - 1][j + 1],
				 	map[i][j - 1],
				 	map[i][j + 1],
				 	map[i + 1][j - 1],
				 	map[i + 1][j],
				 	map[i + 1][j + 1]
				 ].filter(cell => cell === '#').length;
				 
				 switch (map[i][j]) {
				 	case 'L': if (adjacent === 0) { nextMap[i][j] = '#'; hasChange = true; } else { nextMap[i][j] = 'L'; }; break;
				 	case '#': if (adjacent >= 4) { nextMap[i][j] = 'L'; hasChange = true; } else { nextMap[i][j] = '#'; }; break;
				 }
			}
		}

		if (hasChange) {
			yield map;
		} else {
			return map;
		}
		map = nextMap;
	}
}

const iter = gen(map);
while (true) {
	const n = iter.next();
	if (n.done) {
		const result = n.value.map(l => l.filter(l => l === '#').length).reduce((a, b) => a + b, 0);
		console.log(result);
		break;
	}
}
