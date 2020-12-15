const startingNumbers = [5, 2, 8, 16, 18, 0, 1];

const lastSpoken = Object.fromEntries(startingNumbers.map((n, i) => [n, [i]]));
let last = startingNumbers[startingNumbers.length - 1];

for (let i = startingNumbers.length; i < 2020; i++) {
	last = lastSpoken[last].length === 1 ? 0 : lastSpoken[last][0] - lastSpoken[last][1];
	if (!lastSpoken.hasOwnProperty(last)) {
		lastSpoken[last] = [];
	}
	lastSpoken[last].unshift(i);
}

console.log(last);
