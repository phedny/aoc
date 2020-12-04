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

const required = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'];
const optional = ['cid'];

const result = passports
	// Filter passports with only known properties
	.filter(properties => properties.filter(property => !required.includes(property[0]) && !optional.includes(property[0])).length === 0)
	// Filter passports with all required properties
	.filter(properties => required.filter(reqProp => properties.filter(property => property[0] === reqProp).length === 0).length === 0)
	.length;

console.log(result);
