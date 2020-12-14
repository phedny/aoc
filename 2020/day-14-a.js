const fs = require('fs');

function applyMask(value, and, or) {
	// JavaScript does bitwise operators on 32 bits, so let's split the operands first using arithmatic :-(
	const split = v => [v / 0x100000 & 0xffff, v & 0xfffff];
	const [lValue, rValue] = split(value);
	const [lAnd, rAnd] = split(and);
	const [lOr, rOr] = split(or);

	const lResult = lValue & lAnd | lOr;
	const rResult = rValue & rAnd | rOr;
	
	return 0x100000 * lResult + rResult;
}

const machine = { mem: {} };
const list = fs.readFileSync('./inputs/14/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/(mask|mem)(?:\[(\d+)\])? = ([0-9X]+)/))
	.map(([line, maskMem, address, value]) => maskMem === 'mask' ? { line, mask: { or: Number.parseInt(value.replaceAll('X', '0'), 2), and: Number.parseInt(value.replaceAll('X', '1'), 2) } } : { line, assign: { address: Number.parseInt(address), value: Number.parseInt(value) } })
	.forEach(({ line, mask, assign }) => {
		if (mask) machine.mask = mask;
		if (assign) machine.mem[assign.address] = applyMask(assign.value, machine.mask.and, machine.mask.or);
	});

const result = Object.values(machine.mem).reduce((a, b) => a + b, 0);

console.log(result);
