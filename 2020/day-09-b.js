const fs = require('fs');

const numbers = fs.readFileSync('./inputs/9/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => Number.parseInt(line));

function findWrongSum(pLength) {
	for (let i = pLength; i < numbers.length; i++) {
		const expectedSum = numbers[i];
		const hasSome = [...new Array(pLength).keys()].map(n => i - n - 1).some(j => numbers.slice(j, i).includes(expectedSum - numbers[j]));
		if (!hasSome) {
			return expectedSum;
		}
	}
}

const wrongSum = findWrongSum(25);

let start = 0, end = 0, sum;

do {
	sum = numbers.slice(start, end).reduce((a, b) => a + b, 0);
	if (sum > wrongSum) {
		start++;
	} else if (sum < wrongSum) {
		end++;
	}
} while (sum !== wrongSum);

const result = Math.min(...numbers.slice(start, end)) + Math.max(...numbers.slice(start, end));

console.log(result);
