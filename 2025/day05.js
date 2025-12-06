import { getInput } from 'aocparse';

const input = getInput('./day05.in', 'LS-I', 'LI');
const [idRanges, availableIds] = input;

const mergedIdRanges = idRanges.reduce((freshIds, [l, h]) => {
    const overlaps = freshIds.map(([l, h], i) => [l, h, i]).filter(([l1, h1]) => l <= h1 && h >= l1);
    l = Math.min(l, ...overlaps.map((b) => b[0]));
    h = Math.max(h, ...overlaps.map((b) => b[1]));
    const removeIds = overlaps.map((b) => b[2]);
    return [...freshIds.filter((_, i) => !removeIds.includes(i)), [l, h]];
  },
  []);

const freshAvailableIds = availableIds.filter((id) => mergedIdRanges.some(([l, h]) => l <= id && id <= h));
console.log(freshAvailableIds.length);

const freshIdCount = mergedIdRanges.reduce((acc, [l, h]) => acc + h - l + 1, 0);
console.log(freshIdCount);
