import { getInput } from 'aocparse';

const redTiles = getInput('./day09.in', 'LS,I');

const squares = redTiles
  .flatMap(([x1, y1], i) => redTiles.slice(0, i).map(([x2, y2]) => ({ t1: [Math.min(x1, x2), Math.min(y1, y2)], t2: [Math.max(x1, x2), Math.max(y1, y2)], area: (Math.abs(x1-x2)+1) * (Math.abs(y1-y2)+1) })))
  .sort(({ area: a1 }, { area: a2 }) => a2 - a1);

console.log(squares[0].area);

const borderTiles = [...redTiles.slice(1), redTiles[0]]
  .map((t, i) => [redTiles[i], t])
  .flatMap(([[x1, y1], [x2, y2]]) =>
    Array(Math.max(Math.abs(x1-x2), 1)).fill().flatMap((_, i) => x1 + i * Math.sign(x2-x1))
      .flatMap((x) => Array(Math.max(Math.abs(y1-y2), 1)).fill().map((_, i) => [x, y1 + i * Math.sign(y2-y1)]))
  );

borderTiles.splice(0, borderTiles.indexOf(
  borderTiles.reduce((best, curr) => best[0] < curr[0] || (best[0] === curr[0] && best[1] < curr[1]) ? best : curr, [Infinity, Infinity])
)).forEach((t) => borderTiles.push(t));

const outline = [[Infinity, Infinity], [borderTiles[0][0], borderTiles[0][1] - 1]];
borderTiles
  .map((t, i) => [t, borderTiles[(i + 1) % borderTiles.length]])
  .forEach(
    ([[x1, y1], [x2, y2]], i) => {
      const [[x4, y4], [x3, y3]] = outline.slice(outline.length - 2);
      if (x2 === x3 && y2 === y3) {
        outline.pop();
      } else if (y1 === y2 && x1 === x3 && Math.abs(y1 - y3) === 1) {
        outline.push([x2, y3]);
      } else if (x1 === x2 && y1 === y3 && Math.abs(x1 - x3) === 1) {
        outline.push([x3, y2]);
      } else if (y1 === y2 && y1 === y3 && Math.abs(x1 - x2) === 1 && Math.abs(x2 - x3) === 2) {
        outline.push([x3, 2*y3 - y4], [x1, 2*y3 - y4], [x2, 2*y3 - y4]);
      } else if (x1 === x2 && x1 === x3 && Math.abs(y1 - y2) === 1 && Math.abs(y2 - y3) === 2) {
        outline.push([2*x3 - x4, y3], [2*x3 - x4, y1], [2*x3 - x4, y2]);
      } else {
        throw new Error('invalid path');
      }
    },
    [[Infinity, Infinity], [borderTiles[0][0], borderTiles[0][1] - 1]]
  );

const largestGoodSquare = squares
  .filter(({ t1: [x1, y1], t2: [x2, y2] }) => redTiles.every(([x3, y3]) => x3 <= x1 || x3 >= x2 || y3 <= y1 || y3 >= y2))
  .find(({ t1: [x1, y1], t2: [x2, y2] }) => {
  return !outline.some(([x3, y3]) => x1 <= x3 && x3 <= x2 && y1 <= y3 && y3 <= y2)
});

console.log(largestGoodSquare.area);
