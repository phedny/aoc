const fs = require('fs');

const rules = [
	line => line.replace(/\(([0-9 +*]+)\)/, (_, sub) => expr(sub)),
	line => line.replace(/(\d+) \+ (\d+)/, (_, a, b) => Number.parseInt(a) + Number.parseInt(b)),
	line => line.replace(/(\d+) \* (\d+)/, (_, a, b) => Number.parseInt(a) * Number.parseInt(b))
];

function expr(line) {
	for (let rule of rules) {
		while (true) {
			const next = rule(line);
			if (next === line) break;
			line = next;
		}
	}

	return Number.parseInt(line);
}

const result = fs.readFileSync('./inputs/18/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => expr(line))
	.reduce((a, b) => a + b, 0);

console.log(result);
	