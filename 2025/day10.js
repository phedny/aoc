import { Fraction } from 'fraction.js';
import { getInput } from 'aocparse';

const machines = getInput('./day10.in', 'LS ')
  .map((machine) => ({
    lights: machine[0].slice(1, machine[0].length - 1).split('').reduce((acc, c, i) => acc + (c === '#' ? 1 << i : 0), 0),
    buttons: machine.slice(1, machine.length - 1).map((buttons) => buttons.slice(1, buttons.length - 1).split(',').reduce((acc, n) => acc + (1 << parseInt(n)), 0)),
    joltages: machine[machine.length - 1].slice(1, machine[machine.length - 1].length - 1).split(',').map((n) => parseInt(n)),
  }));

const minPushes = machines.map(({ lights, buttons }) =>
  Array(1 << buttons.length)
    .fill()
    .map((_, i) => Array(buttons.length).fill().map((_, i) => i).filter((j) => i & (1<<j)).map((i) => buttons[i]))
    .map((is) => [is.reduce((acc, n) => acc ^ n, 0), is.length])
    .filter(([s]) => s === lights)
    .reduce((acc, [_, c]) => Math.min(acc, c), Infinity)
);

console.log(minPushes.reduce((a, b) => a + b, 0));

const results = machines
  .map(({ joltages, buttons }) => {
    const matrix = joltages.map((count, i) => [...buttons.map((b) => b & (1<<i) ? Fraction(1) : Fraction(0)), Fraction(count)]);
    let fixedRows = 0;
    for (let col = 0; col < matrix[0].length - 1; col++) {
      const nonZero = matrix.flatMap((r, i) => r[col].equals(0) ? [] : [i]);
      const hitIdx = nonZero.findIndex((n) => n >= fixedRows);
      if (hitIdx > -1) {
        const [hit] = nonZero.splice(hitIdx, 1);
        for (const j of nonZero) {
          const factor = matrix[j][col].div(matrix[hit][col]);
          matrix[j] = matrix[j].map((v, i) => v.sub(factor.mul(matrix[hit][i])));
        }
        const [row] = matrix.splice(hit, 1);
        matrix.splice(fixedRows++, 0, row.map((v) => v.div(row[col])));
      }
    }
    for (let i = 0; i < matrix.length; ) {
      if (matrix[i].some((v) => !v.equals(0))) {
        i++;
      } else {
        matrix.splice(i, 1);
      }
    }
    const numberOfPresses = new Map();
    let numberOfPressesFound;
    do {
      numberOfPressesFound = false;
      for (let i = 0; i < matrix.length; ) {
        const nonZero = matrix[i].slice(0, matrix[i].length - 1).flatMap((c, i) => c.equals(0) ? [] : [i]);
        if (nonZero.length === 1) {
          const button = nonZero[0];
          const pressCount = matrix[i][matrix[i].length - 1].div(matrix[i][nonZero[0]])
          numberOfPresses.set(button, pressCount.valueOf());
          matrix.splice(i, 1);
          for (const row of matrix) {
            row[row.length - 1] = row[row.length - 1].sub(row[button].mul(pressCount))
            row[button] = Fraction(0);
          }
          numberOfPressesFound = true
        } else {
          i++;
        }
      }
    } while (numberOfPressesFound);
    let result = [...numberOfPresses.values()].reduce((a, b) => a + b, 0)
    if (matrix.length > 0) {
      const computeCols = matrix[0].slice(0, matrix[0].length).flatMap((_, i) => matrix.filter((r) => !r[i].equals(0)).length === 1 ? [i] : [])
      const tweakCols = matrix[0].slice(0, matrix[0].length - 1)
        .map((_, i) => i)
        .filter((i) => matrix.some((r) => !r[i].equals(0)) && !computeCols.includes(i))
        .slice(0, matrix[0].length - matrix.length - 1);

      const maxPushes = buttons.map((b) => joltages.map((joltage, i) => b & (1<<i) ? joltage : Infinity).reduce((a, b) => Math.min(a, b), Infinity))
      result += tweakCols
        .reduce((acc, i) => acc.flatMap((acc) => Array(maxPushes[i] + 1).fill().map((_, i) => [...acc, i])), [[]])
        .map((inputs) => {
          const rows = [...matrix];
          const known = new Map(inputs.map((v, i) => [tweakCols[i], Fraction(v)]));
          while (known.size < computeCols.length + tweakCols.length) {
            const oldLen = rows.length;
            for (let i = 0; i < rows.length; ) {
              const row = rows[i];
              const [sum, unknown] = row.slice(0, row.length - 1).reduce(([sum, unknown], c, i) => known.has(i) ? [sum.add(c.mul(known.get(i))), unknown] : c.equals(0) ? [sum, unknown] : [sum, [...unknown, i]], [Fraction(0), []])
              if (unknown.length === 1) {
                known.set(unknown[0], row[row.length - 1].sub(sum).div(row[unknown[0]]));
                rows.splice(i, 1);
              } else {
                i++;
              }
            }
            if (oldLen === rows.length) {
              console.log('Whut!', known, rows.map((r) => r.join(', ')));
              process.exit();
            }
          }
          const pressCount = [...known.values()];
          return pressCount.every((c) => c.divisible(1) && c.gte(0)) ? pressCount.reduce((a, b) => a + b.valueOf(), 0) : Infinity;
        })
        .reduce((a, b) => a < b ? a : b, [Infinity]);
    }
    return result;
  });

console.log(results.reduce((a, b) => a + b, 0));
