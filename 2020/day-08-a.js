const fs = require('fs');

const program = fs.readFileSync('./inputs/8/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/^(\w+) ([+-]\d+)$/))
	.map(([_, instr, arg]) => ({ instr, arg }));;

let acc = 0;
const visited = new Set();

for (let ip = 0; !visited.has(ip); ip++) {
	visited.add(ip);
	const { instr, arg } = program[ip];
	const argN = Number.parseInt(arg);
	switch (instr) {
		case 'jmp': ip += argN - 1; break;
		case 'acc': acc += argN; break;
		case 'nop': break;
	}
}

console.log(acc);
