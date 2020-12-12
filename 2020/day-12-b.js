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
			case 'L': return [0, 0, 0, 360 - num];
			case 'F': return [0, 0, num, 0];
		}
	})
	.reduce(([n, e, wN, wE], [dN, dE, f, wR]) => {
		let tN, tE;
		switch (wR) {
			case 0:   [tN, tE] = [wN, wE]; break;
			case 90:  [tN, tE] = [-wE, wN]; break;
			case 180: [tN, tE] = [-wN, -wE]; break;
			case 270: [tN, tE] = [wE, -wN]; break;
		}
		return [n + f * wN, e + f * wE, tN + dN, tE + dE];
	}, [0, 0, 1, 10]);

const result = Math.abs(end[0]) + Math.abs(end[1]);

console.log(result);
