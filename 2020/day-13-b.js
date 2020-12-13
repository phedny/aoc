const fs = require('fs');

const [_, busList] = fs.readFileSync('./inputs/13/real.txt', { encoding: 'utf8' }).split('\n');
const list = busList.split(',')
	.map((line, i) => line === 'x' ? [0, 1] : [(i * line - i) % line, line]);

const result = list.reduce(([v1, m1], [v2, m2]) => {
		for (let vR = v1 + m1; ; vR += m1) {
			if (vR % m2 === v2) {
				return [vR % (m1 * m2), m1 * m2];
			}
		}
	}, [0, 1]);

console.log(result[0]);
