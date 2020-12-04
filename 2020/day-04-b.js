const fs = require('fs');

const passports = [];
var properties = [];

fs.readFileSync('./inputs/4/real.txt', { encoding: 'utf8' })
	.split('\n')
	.forEach(readLine);
readLine();

function readLine(line) {
	if (line && line.length > 0) {
		return properties.splice(properties.length, 0, ...line.split(' ').map(item => item.split(':')));
	}

	passports.push(properties);
	properties = [];
}

const required = {
	byr: v => {
		const i = Number.parseInt(v);
		return i.toString() === v && i >= 1920 && i <= 2002;
	},
	iyr: v => {
		const i = Number.parseInt(v);
		return i.toString() === v && i >= 2010 && i <= 2020;
	},
	eyr: v => {
		const i = Number.parseInt(v);
		return i.toString() === v && i >= 2020 && i <= 2030;
	},
	hgt: v => {
		const m = v.match(/^(\d+)(in|cm)$/);
		if (!m) return false;
		const [_, n, u] = m;
		const i = Number.parseInt(n);
		switch (u) {
			case 'cm': return i.toString() === n && i >= 150 && i <= 193;
			case 'in': return i.toString() === n && i >= 59 && i <= 76;
			default: return false;
		}
	},
	hcl: v => !!(v.match(/^#[0-9a-f]{6}$/)),
	ecl: v => ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].includes(v),
	pid: v => !!(v.match(/^[0-9]{9}$/))
};
const optional = ['cid'];

const result = passports
	// Filter passports with only known properties
	.filter(properties => properties.filter(property => !Object.keys(required).includes(property[0]) && !optional.includes(property[0])).length === 0)
	// Filter passports with all required properties
	.filter(properties => Object.entries(required).filter(([reqProp, test]) => properties.filter(property => property[0] === reqProp && test(property[1])).length === 0).length === 0)
	.length;

console.log(result);
