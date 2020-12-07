const fs = require('fs');

const bags = fs.readFileSync('./inputs/7/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/(\w+ \w+) bags contain (.+)\./))
	.map(([_, containingBag, containedBags]) => [containingBag, containedBags.split(', ').map(line => line.match(/(\d+) (\w+ \w+) bag/))])
	.map(([containingBag, containedBags]) => [containingBag, containedBags[0] ? containedBags : []])
	.map(([containingBag, containedBags]) => [containingBag, containedBags.map(([_, num, type]) => [Number.parseInt(num), type])])
	.flatMap(([containingBag, containedBags]) => containedBags.map(([_, type]) => [type, containingBag]));

const map = new Map(bags.map(([containingBag]) => [containingBag, []]));
bags.forEach(([containingBag, containedBag]) => map.get(containingBag).push(containedBag));

const curr = new Set(map.get('shiny gold'));
let next;

do {
	next = false;
	Array.from(curr).flatMap(a => map.get(a))
		.forEach(a => {
			if (a && !curr.has(a)) {
				curr.add(a);
				next = true;
			}
		});
} while(next);

console.log(curr.size);
