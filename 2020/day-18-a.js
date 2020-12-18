const fs = require('fs');

function apply(operator, op1, op2) {
	if (operator === '+') {
		return op1 + op2;
	} else if (operator === '*') {
		return op1 * op2;
	} else {
		return op2;
	}
}

function expr(line) {
	const stack = [];
	let num, operator;

	while (line.length > 0) {
		const char = line.shift();
		if (char === ' ') continue;

		if (char >= '0' && char <= '9') {
			num = apply(operator, num, Number.parseInt(char));
			operator = null;
		}

		if (char === '+' || char === '*') {
			operator = char;
		}

		if (char === '(') {
			stack.push([num, operator]);
			num = null;
			operator = null;
		}

		if (char === ')') {
			const [sNum, sOperator] = stack.pop();
			num = apply(sOperator, sNum, num);
			operator = null;
		}
	}

	return num;
}

const result = fs.readFileSync('./inputs/18/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => expr(line.split('')))
	.reduce((a, b) => a + b, 0);

console.log(result);
	