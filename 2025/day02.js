const testableIds = require('fs')
  .readFileSync('./day02.in', { encoding: 'utf-8' })
  .split(',')
  .map((range) => range.split('-').map((n) => parseInt(n)))
  .flatMap(([low, high]) => Array(high - low + 1).fill().map((_, n) => low + n))
  .map((n) => [String(n), n]);

const countA = testableIds
  .filter(([n]) => n.substring(0, n.length / 2) === n.substring(n.length / 2))
  .reduce((sum, [_, n]) => sum + n, 0)
console.log(countA);

const countB = testableIds
  .filter(([n]) => Array(n.length).fill().map((_, i) => i).filter((i) => n.length % i === 0).some((l) => {
      const ps = Array(n.length / l).fill().map((_, i) => n.slice(l * i, l * (i + 1)));
      return ps.every((p) => p === ps[0]);
    }))
  .reduce((sum, [_, n]) => sum + n, 0)
console.log(countB);
