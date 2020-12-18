const fs = require('fs');

const input = (function* () { for (const line of fs.readFileSync('./inputs/16/real.txt', { encoding: 'utf8' }).split('\n')) yield line; })();

function* propsFrom(input) {
	while (true) {
		const { value: line, done } = input.next();
		if (done || line === '') return;
		
		const match = line.match(/([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)/);
		yield {
			name: match[1],
			ranges: [
				[Number.parseInt(match[2]), Number.parseInt(match[3])],
				[Number.parseInt(match[4]), Number.parseInt(match[5])]
			]
		};
	}
}

const props = [];
for (const prop of propsFrom(input)) props.push(prop);

input.next(); // your ticket
const myTicket = input.next().value.split(',').map(a => Number.parseInt(a));

input.next(); // empty line
input.next(); // nearby tickets

const nearbyTickets = [];
for (const line of input) nearbyTickets.push(line.split(',').map(a => Number.parseInt(a)));

const allRanges = props.flatMap(prop => prop.ranges);

const possibleFields = nearbyTickets
	.filter(ticket => ticket.every(field => allRanges.some(range => field >= range[0] && field <= range[1])))
	.map(ticket => ticket.map(field => props.filter(prop => prop.ranges.some(range => field >= range[0] && field <= range[1])).map(prop => prop.name)))
	.reduce((a, b) => a.map((aField, i) => aField.filter(p => b[i].includes(p))), myTicket.map(() => props.map(prop => prop.name)))
	.map(p => ({ possible: p }));

while (possibleFields.some(field => !field.hasOwnProperty('mustBe'))) {
	const foundFields = possibleFields.filter(field => field.possible.length === 1)
		.map(field => field.mustBe = field.possible[0]);
	possibleFields.forEach(field => field.possible = field.possible.filter(p => !foundFields.includes(p)));
}

const result = possibleFields.map((field, i) => ({ field: field.mustBe, value: myTicket[i] }))
	.filter(({ field }) => field.startsWith('departure'))
	.map(({ value }) => value)
	.reduce((a, b) => a * b, 1);

console.log(result);
