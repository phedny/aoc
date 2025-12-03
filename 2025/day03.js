const banks = require('fs')
  .readFileSync('./day03.in', { encoding: 'utf-8' })
  .split('\n');

const sumA = banks
  .map((bank) => parseInt(highestJoltage(bank, 2)))
  .reduce((a, b) => a + b, 0);

console.log(sumA);

const sumB = banks
  .map((bank) => parseInt(highestJoltage(bank, 12)))
  .reduce((a, b) => a + b, 0);

console.log(sumB);

function highestJoltage(bank, n) {
  if (n === 0) {
    return '';
  }
  const i = bank.substring(0, bank.length - n + 1).split('').sort((a, b) => b - a)[0];
  return i + highestJoltage(bank.substring(bank.indexOf(i) + 1), n - 1);
}
