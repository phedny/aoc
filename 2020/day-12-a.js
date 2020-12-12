const fs = require('fs');

const end = fs.readFileSync('./inputs/12/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => {
		const num = Number.parseInt(line.substring(1));
		switch (line.substring(0, 1)) {
			case 'N': return [num, 0, 0, 0];
			case 'E': return [0, num, 0, 0];
			case 'S': return [-num, 0, 0, 0];
			case 'W': return [0, -num, 0, 0];
			case 'R': return [0, 0, 0, num];
			case 'L': return [0, 0, 0, -num];
			case 'F': return [0, 0, num, 0];
		}
	})
	.reduce(([n, e, o], [dN, dE, f, dO]) => {
		const dN2 = o === 0 ? f : o === 180 ? -f : 0;
		const dE2 = o === 90 ? f : o === 270 ? -f : 0;
		return [n + dN + dN2, e + dE + dE2, (360 + o + dO) % 360];
	}, [0, 0, 90]);

const result = Math.abs(end[0]) + Math.abs(end[1]);

console.log(result);
