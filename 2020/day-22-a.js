const fs = require('fs');

const cardsPerPlayer = fs.readFileSync('./inputs/22/real.txt', { encoding: 'utf8' })
	.split('\n\n')
	.map(playerLines => playerLines.split('\n'))
	.map(lines => { lines.shift(); return lines.map(line => Number.parseInt(line)); });

while (cardsPerPlayer.every(cards => cards.length > 0)) {
	const roundCards = cardsPerPlayer.map(cards => cards.shift());
	const maxCard = Math.max(...roundCards);
	const winningPlayer = roundCards.map((card, i) => [card, i])
		.filter(([card]) => card === maxCard)
		.map(([_, i]) => i);
	cardsPerPlayer[winningPlayer].push(...roundCards.sort((a, b) => b - a));
}

const result = cardsPerPlayer.filter(cards => cards.length > 0)[0]
	.reverse()
	.reduce((a, b, i) => a + b * (i + 1), 0);

console.log(result);
