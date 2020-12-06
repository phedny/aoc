const fs = require('fs');

const groups = [];
let answers = '';

fs.readFileSync('./inputs/6/real.txt', { encoding: 'utf8' })
	.split('\n')
	.forEach(readLine);
readLine();

function readLine(line) {
	if (line && line.length > 0) {
		return answers += line;
	}

	groups.push(answers);
	answers = '';
}

const result = groups.map(answers => new Set(answers.split('')).size)
	.reduce((a, b) => a + b, 0);

console.log(result);
