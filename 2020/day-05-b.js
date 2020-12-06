const fs = require('fs');

const numbers = fs.readFileSync('./inputs/5/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.split('').map(c => ['B', 'R'].includes(c) ? 1 : 0).reverse().map((d, i) => d * Math.pow(2, i)).reduce((a, b) => a + b, 0));

let firstFullRow = 0;
while ([...new Array(8).keys()].map((_, i) => 8 * firstFullRow + i).some(i => !numbers.includes(i))) firstFullRow++;

let highestSeat = Math.max(...numbers);
let lastFullRow = Math.floor(highestSeat / 8);
while ([...new Array(8).keys()].map((_, i) => 8 * lastFullRow + i).some(i => !numbers.includes(i))) lastFullRow--;

const result = numbers.filter(n => n >= 8 * firstFullRow && n < 8 * (lastFullRow + 1)).reduce((a, b) => a ^ b, 0);

console.log(result);
