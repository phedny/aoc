import { getInput } from 'aocparse';

const grid = getInput('./day07.in', 'LC');

function count(multiworld) {
  const [splits] = grid.slice(1).reduce(
    ([splits, beams], line) => {
      const nextBeams = {};
      const parsed = Object.entries(beams)
        .map(([x, count]) => [parseInt(x), count]);
      parsed.flatMap(([x, count]) => line[x] === '^' ? [[x - 1, count], [x + 1, count]] : [[x, count]])
        .forEach(([x, count]) => { nextBeams[x] = (nextBeams[x] ?? 0) + count });
      return [parsed.reduce(((splits, [x, count]) => line[x] === '^' ? splits + (multiworld ? count : 1) : splits), splits), nextBeams];
    },
    [multiworld ? 1 : 0, { [grid[0].indexOf('S')]: 1 }],
  );
  return splits;
}

console.log(count(false));
console.log(count(true));
