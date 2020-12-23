// let cups = '389125467'.split('').map(num => Number.parseInt(num));
let cups = '135468729'.split('').map(num => Number.parseInt(num));

for (let i = 0; i < 100; i++) {
	const next = cups.indexOf([1, 2, 3].map(n => cups[0] - n).filter(n => !cups.slice(1, 4).includes(n))[0] || [0, 1, 2].map(n => cups.length - n).filter(n => !cups.slice(1, 4).includes(n))[0]);
	cups = [...cups.slice(4, next + 1), ...cups.slice(1, 4), ...cups.slice(next + 1), cups[0]];
}

const indexOfOne = cups.indexOf(1);
const result = [...cups.slice(indexOfOne + 1), ...cups.slice(0, indexOfOne)];

console.log(result.join(''));
