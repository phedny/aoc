const lines = require('fs')
  .readFileSync('./day06.in', { encoding: 'utf-8' })
  .split('\n');

const [lastLine] = lines.splice(lines.length - 1);
const operators = lastLine.trim().split(/ +/);

const results = lines.reduce(
  (acc, line) => {
    const ns = line.trim().split(/ +/).map((n) => parseInt(n));
    return acc.map((v, i) => operators[i] === '*' ? ns[i] * v : ns[i] + v);
  },
  operators.map((op) => op === '*' ? 1 : 0),
);
console.log(results.reduce((acc, n) => acc + n, 0));

const cols = lastLine.split('').map((c, i) => [parseInt(lines.map((line) => line[i]).join('').trim()), c]);
cols.push([NaN, ' ']);

const [sum] = cols.reduce(
  ([sum, acc, op], [n, c]) => !acc ? [sum, n, c] : isNaN(n) ? [sum + acc] : [sum, op === '*' ? acc * n : acc + n, op],
  [0],
)
console.log(sum);
