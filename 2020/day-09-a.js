const fs = require('fs');

const numbers = fs.readFileSync('./inputs/9/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => Number.parseInt(line));

function findWeakness(pLength) {
	for (let i = pLength; i < numbers.length; i++) {
		const expectedSum = numbers[i];
		const hasSome = [...new Array(pLength).keys()].map(n => i - n - 1).some(j => numbers.slice(j, i).includes(expectedSum - numbers[j]));
		if (!hasSome) {
			return expectedSum;
		}
	}
}

const result = findWeakness(25);

console.log(result);
