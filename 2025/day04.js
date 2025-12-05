const rolls = require('fs')
  .readFileSync('./day04.in', { encoding: 'utf-8' })
  .split('\n')
  .map((line) => line.split('').map((c) => c === '@'));

function neighbors8(x, y) {
  return [
    [x - 1, y - 1], [x - 1, y], [x - 1, y + 1],
    [x, y - 1], [x, y + 1],
    [x + 1, y - 1], [x + 1, y], [x + 1, y + 1],
  ];
}

function countRolls(rolls) {
  return rolls.reduce((acc, line) => acc + line.filter((c) => c).length, 0);
}

function hasRoll(rolls, ...cs) {
  for (const c of cs) {
    if (c < 0 || c >= rolls.length) {
      return false;
    }
    rolls = rolls[c];
  }
  return rolls;
}

function findAccessibleRolls(rolls) {
  return rolls.flatMap((line, x) =>
    line.flatMap((r, y) =>
      r && neighbors8(x, y).filter(([x, y]) => hasRoll(rolls, x, y)).length < 4 ? [[x, y]] : []
    )
  );
}

function deleteRolls(rolls, cs) {
  const clone = rolls.map((line) => [...line]);
  for (const c of cs) {
    clone[c[0]][c[1]] = false;
  }
  return clone;
}

console.log(findAccessibleRolls(rolls).length);

let currRolls = rolls;
while (true) {
  const accessibleRolls = findAccessibleRolls(currRolls);
  if (accessibleRolls.length === 0) {
    break;
  }
  currRolls = deleteRolls(currRolls, accessibleRolls);
}

console.log(countRolls(rolls) - countRolls(currRolls));
