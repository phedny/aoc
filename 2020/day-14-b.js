const fs = require('fs');

function expand(value) {
	return value.split('').reduce((acc, l) => l === 'X' ? [...acc.map(a => a + '0'), ...acc.map(a => a + '1')] : acc.map(a => a + l), ['']);
}

function applyMask(value, and, or, float) {
	// JavaScript does bitwise operators on 32 bits, so let's split the operands first using arithmatic :-(
	const split = v => [v / 0x100000 & 0xffff, v & 0xfffff];
	const [lValue, rValue] = split(value);
	const [lAnd, rAnd] = split(and);
	const [lOr, rOr] = split(or);
	const [lFloat, rFloat] = split(float);

	const lResult = lValue & lAnd | lOr | lFloat;
	const rResult = rValue & rAnd | rOr | rFloat;
	
	return 0x100000 * lResult + rResult;
}

const machine = { mem: {} };
const list = fs.readFileSync('./inputs/14/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/(mask|mem)(?:\[(\d+)\])? = ([0-9X]+)/))
	.map(([line, maskMem, address, value]) => maskMem === 'mask' ? { line, mask: { or: Number.parseInt(value.replaceAll('X', '0'), 2), and: Number.parseInt(value.replaceAll('0', '1').replaceAll('X', '0'), 2), float: expand(value).map(float => Number.parseInt(float, 2)) } } : { line, assign: { address: Number.parseInt(address), value: Number.parseInt(value) } })
	.forEach(({ line, mask, assign }) => {
		if (mask) machine.mask = mask;
		if (assign) machine.mask.float.map(float => applyMask(assign.address, machine.mask.and, machine.mask.or, float)).forEach(address => machine.mem[address] = assign.value);
	});

const result = Object.values(machine.mem).reduce((a, b) => a + b, 0);

console.log(result);
