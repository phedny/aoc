const fs = require('fs');

const input = (function* () { for (const line of fs.readFileSync('./inputs/19/real.txt', { encoding: 'utf8' }).split('\n')) yield line; })();

const rules = {};

while (true) {
	const { value: line, done } = input.next();
	if (done || line === '') break;

	let match;
	if (match = line.match(/^(\d+): "(\w)"$/)) {
		rules[match[1]] = { literal: match[2] };
	} else if (match = line.match(/^(\d+): ([0-9 ]+)$/)) {
		rules[match[1]] = { sequence: match[2].split(' ') };
	} else if (match = line.match(/^(\d+): ([0-9 ]+) \| ([0-9 ]+)$/)) {
		rules[match[1]] = { or: [{ sequence: match[2].split(' ') }, { sequence: match[3].split(' ') }] };
	} else {
		console.log('No match', line);
	}
}

Object.values(rules).forEach(rule => {
	if (rule.hasOwnProperty('literal')) {
	} else if (rule.hasOwnProperty('sequence')) {
		rule.validate = (s) => rule.sequence.map(seq => rules[seq]).reduce((acc, seq) => acc, [s]);
		rule.sequence = rule.sequence.map(seq => rules[seq]);
	} else if (rule.hasOwnProperty('or')) {
		rule.or = rule.or.map(or => ({ sequence: or.sequence.map(seq => rules[seq]) }));
	}
});

function validate(input, rule) {
	if (rule.hasOwnProperty('literal')) {
		return input.length > 0 && input.charAt(0) === rule.literal ? [input.substring(1)] : [];
	} else if (rule.hasOwnProperty('sequence')) {
		return rule.sequence.reduce((acc, seq) => acc.flatMap(i => validate(i, seq)), [input]);
	} else if (rule.hasOwnProperty('or')) {
		return rule.or.flatMap(r => validate(input, r));
	}
}

let result = 0;
for (const line of input) {
	if (validate(line, rules[0]).some(r => r.length === 0)) result++;
}

console.log(result);
