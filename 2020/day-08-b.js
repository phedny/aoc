const fs = require('fs');

const originalProgram = fs.readFileSync('./inputs/8/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/^(\w+) ([+-]\d+)$/))
	.map(([_, instr, arg]) => ({ instr, arg }));;

function run(program) {
	let acc = 0;
	const visited = new Set();

	for (let ip = 0; ip < program.length; ip++) {
		if (visited.has(ip)) {
			return null;
		}

		visited.add(ip);
		const { instr, arg } = program[ip];
		const argN = Number.parseInt(arg);
		switch (instr) {
			case 'jmp': ip += argN - 1; break;
			case 'acc': acc += argN; break;
			case 'noop': break;
		}
	}

	return acc;
}

let result;

for (let i = 0; i < originalProgram.length; i++) {
	if (originalProgram[i].instr === 'acc') continue;

	const originalInstr = originalProgram[i].instr;
	switch (originalInstr) {
		case 'noop': originalProgram[i].instr = 'jmp'; break;
		case 'jmp': originalProgram[i].instr = 'noop'; break;
	}

	result = run(originalProgram);
	originalProgram[i].instr = originalInstr;

	if (result) break;
}

console.log(result);
