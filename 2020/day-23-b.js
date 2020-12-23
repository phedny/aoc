// let firstCups = '389125467'.split('').map(num => ({ num: Number.parseInt(num) }));
let firstCups = '135468729'.split('').map(num => ({ num: Number.parseInt(num) }));

let cupOne, cupNine;
firstCups.forEach((cup, i) => {
	cup.prev = firstCups[(i + 8) % 9];
	cup.next = firstCups[(i + 1) % 9];
	cup.minusOne = firstCups.filter(cup2 => cup2.num === cup.num - 1)[0];
	if (cup.num === 1) cupOne = cup;
	if (cup.num === 9) cupNine = cup;
});
cupOne.minusOne = cupNine;

let currentCup = firstCups[0];
let prevCup = cupNine;
for (let i = 10; i <= 1000000; i++) {
	const newCup = {
		num: i,
		prev: prevCup,
		minusOne: prevCup
	};

	if (i === 10) {
		firstCups[8].next = newCup;
	} else {
		prevCup.next = newCup;
	}

	prevCup = newCup;
}

prevCup.next = currentCup;
currentCup.prev = prevCup;

cupOne.minusOne = prevCup;

for (let i = 0; i < 10000000; i++) {
	const moveCups = [currentCup.next, currentCup.next.next, currentCup.next.next.next];
	let behindCup = currentCup.minusOne;
	while (moveCups.includes(behindCup)) behindCup = behindCup.minusOne;

	// Do the move
	[currentCup.next, behindCup.next, moveCups[2].next, moveCups[2].next.prev, moveCups[0].prev, behindCup.next.prev] = [moveCups[2].next, moveCups[0], behindCup.next, currentCup, behindCup, moveCups[2]];
	currentCup = currentCup.next;
}

console.log(cupOne.next.num * cupOne.next.next.num);
