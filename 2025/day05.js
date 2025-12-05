const input = require('fs')
  .readFileSync('./day05.in', { encoding: 'utf-8' })
  .split('\n\n');

const freshIds = input[0]
  .split('\n')
  .map((line) => line.split('-').map((n) => parseInt(n)))
  .reduce((freshIds, [l, h]) => {
    const overlaps = freshIds.map(([l, h], i) => [l, h, i]).filter(([l1, h1]) => l <= h1 && h >= l1);
    l = Math.min(l, ...overlaps.map((b) => b[0]));
    h = Math.max(h, ...overlaps.map((b) => b[1]));
    const removeIds = overlaps.map((b) => b[2]);
    return [...freshIds.filter((_, i) => !removeIds.includes(i)), [l, h]];
  },
  []);

const availableIds = input[1]
  .split('\n')
  .map((n) => parseInt(n));

const freshAvailableIds = availableIds.filter((id) => freshIds.some(([l, h]) => l <= id && id <= h));
console.log(freshAvailableIds.length);

const freshIdCount = freshIds.reduce((acc, [l, h]) => acc + h - l + 1, 0);
console.log(freshIdCount);
