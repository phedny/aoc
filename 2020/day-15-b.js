const startingNumbers = [5, 2, 8, 16, 18, 0, 1];

const lastSpoken = new Array(30000000);
lastSpoken[29999999] = null;
startingNumbers.forEach((n, i) => lastSpoken[n] = [i]);
let last = startingNumbers[startingNumbers.length - 1];

for (let i = startingNumbers.length; i < 30000000; i++) {
	last = lastSpoken[last].length === 1 ? 0 : lastSpoken[last][0] - lastSpoken[last][1];
	if (!lastSpoken[last]) {
		lastSpoken[last] = [i];
	} else {
		lastSpoken[last].length = 2;
		lastSpoken[last][1] = lastSpoken[last][0];
		lastSpoken[last][0] = i;
	}
}

console.log(last);
