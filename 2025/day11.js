import { getInput } from 'aocparse';

const wires = getInput('./day11.in', 'LS:')
    .map(([from, to]) => [from, to.substring(1).split(' ')]);

function countPaths(start, end, forbidden) {
  let myWires = wires
    .filter(([from]) => from !== end && !forbidden.includes(from))
    .map(([from, to]) => [from, to.filter((v) => !forbidden.includes(v))])
    .filter(([_, to]) => to.length);
  let count = Infinity;
  while (myWires.reduce((acc, [_, to]) => acc + to.length, 0) !== count) {
    count = myWires.reduce((acc, [_, to]) => acc + to.length, 0);
    myWires = myWires
      .filter(([from]) => from === start || myWires.some(([_, to]) => to.includes(from)))
      .map(([from, to]) => [from, to.filter((v) => v === end || myWires.some(([from]) => v === from))])
      .filter(([_, to]) => to.length);
  }
  const todo = new Map(myWires.map(([from, to]) => [from, new Set(to)]));
  const pathsToEnd = new Map([[end, 1]]);
  let more = true;
  while (more) {
    more = false;
    for (const [from, to] of todo) {
      for (const n of [...to.values()].filter((v) => pathsToEnd.has(v) && !todo.has(v))) {
        pathsToEnd.set(from, (pathsToEnd.get(from) ?? 0) + pathsToEnd.get(n));
        to.delete(n);
        more = true;
      }
      if (!to.size) {
        todo.delete(from);
      }
    }
  }
  return pathsToEnd.get(start) ?? 0;
}

const youOut = countPaths('you', 'out', []);
console.log(youOut);

const svrDacFftOut = countPaths('svr', 'dac', ['fft', 'out']) * countPaths('dac', 'fft', ['svr', 'out']) * countPaths('fft', 'out', ['svr', 'dac']);
const svrFftDacOut = countPaths('svr', 'fft', ['dac', 'out']) * countPaths('fft', 'dac', ['svr', 'out']) * countPaths('dac', 'out', ['svr', 'fft']);
console.log(svrDacFftOut + svrFftDacOut);
