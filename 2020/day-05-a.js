const fs = require('fs');

const numbers = fs.readFileSync('./inputs/5/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.split('').map(c => ['B', 'R'].includes(c) ? 1 : 0).reverse().map((d, i) => d * Math.pow(2, i)).reduce((a, b) => a + b, 0));

const result = Math.max(...numbers);

console.log(result);
