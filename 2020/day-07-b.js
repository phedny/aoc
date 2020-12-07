const fs = require('fs');

const bags = Object.fromEntries(
	fs.readFileSync('./inputs/7/real.txt', { encoding: 'utf8' })
		.split('\n')
		.map(line => line.match(/(\w+ \w+) bags contain (.+)\./))
		.map(([_, containingBag, containedBags]) => [containingBag, containedBags.split(', ').map(line => line.match(/(\d+) (\w+ \w+) bag/))])
		.map(([containingBag, containedBags]) => [containingBag, containedBags[0] ? containedBags : []])
		.map(([containingBag, containedBags]) => [containingBag, containedBags.map(([_, num, type]) => [Number.parseInt(num), type])])
);

let result = 0;
let curr = [[1, 'shiny gold']];

do {
	curr = curr.flatMap(([num1, type1]) => bags[type1].map(([num2, type2]) => [num1 * num2, type2]));
	result += curr.map(([num1]) => num1).reduce((a, b) => a + b, 0);
} while(curr.length);

console.log(result);
