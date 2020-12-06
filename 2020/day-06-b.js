const fs = require('fs');

const groups = [];
let answers = new Set([...new Array(26).keys()].map((_, i) => String.fromCharCode(97 + i)));

fs.readFileSync('./inputs/6/real.txt', { encoding: 'utf8' })
	.split('\n')
	.forEach(readLine);
readLine();

function readLine(line) {
	if (line && line.length > 0) {
		return [...answers].filter(answer => line.indexOf(answer) === -1)
			.forEach(answer => answers.delete(answer));
	}

	groups.push(answers);
	answers = new Set([...new Array(26).keys()].map((_, i) => String.fromCharCode(97 + i)));
}

const result = groups.map(answers => answers.size)
	.reduce((a, b) => a + b, 0);

console.log(result);
