const [countA, countB] = require('fs')
  .readFileSync('./day01.in', { encoding: 'utf-8' })
  .split('\n')
  .flatMap((line) => Array(parseInt(line.substring(1))).fill(line[0]).concat(['A']))
  .reduce(
    ([countA, countB, pos], step) => {
      switch (step) {
        case 'L':
          switch (pos) {
            case 0:
              return [countA, countB, 99];
            case 1:
              return [countA, countB + 1, 0];
            default:
              return [countA, countB, pos - 1];
          }
        case 'R':
          switch (pos) {
            case 99:
              return [countA, countB + 1, 0];
            default:
              return [countA, countB, pos + 1];
          }
        case 'A':
          switch (pos) {
            case 0:
              return [countA + 1, countB, pos];
            default:
              return [countA, countB, pos];
          }
      }
    },
    [0, 0, 50]
  );
console.log(countA);
console.log(countB);
