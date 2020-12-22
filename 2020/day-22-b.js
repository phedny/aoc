const fs = require('fs');

const cardsPerPlayer = fs.readFileSync('./inputs/22/real.txt', { encoding: 'utf8' })
	.split('\n\n')
	.map(playerLines => playerLines.split('\n'))
	.map(lines => { lines.shift(); return lines.map(line => Number.parseInt(line)); });

function recursiveCombat(cardsPerPlayer) {
	const seenHands = [];
	while (cardsPerPlayer.every(cards => cards.length > 0)) {
		const handsAsString = cardsPerPlayer.map(cards => cards.join(',')).join('-');
		if (seenHands.includes(handsAsString)) {
			return [0, cardsPerPlayer[0]];
		}
		seenHands.push(handsAsString);
		const roundCards = cardsPerPlayer.map(cards => cards.shift());
		if (roundCards.every((card, i) => cardsPerPlayer[i].length >= card)) {
			const [winningPlayer] = recursiveCombat(cardsPerPlayer.map((cards, i) => cards.slice(0, roundCards[i])));
			cardsPerPlayer[winningPlayer].push(roundCards[winningPlayer], roundCards[1 - winningPlayer]);
		} else {
			const maxCard = Math.max(...roundCards);
			const winningPlayer = roundCards.map((card, i) => [card, i])
				.filter(([card]) => card === maxCard)
				.map(([_, i]) => i);
			cardsPerPlayer[winningPlayer].push(...roundCards.sort((a, b) => b - a));
		}
	}

	const winningPlayer = cardsPerPlayer.map((cards, i) => [cards, i])
		.filter(([cards]) => cards.length > 0)
		.map(([_, i]) => i);
	return [winningPlayer, cardsPerPlayer[winningPlayer]];
}

const [winningPlayer, winningHand] = recursiveCombat(cardsPerPlayer);
const result = winningHand.reverse()
	.reduce((a, b, i) => a + b * (i + 1), 0);

console.log(result);
