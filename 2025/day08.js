import { getInput } from 'aocparse';

const junctionBoxes = getInput('./day08.in', 'LS,I');

const distances = junctionBoxes
  .flatMap(([x1, y1, z1], n1) => junctionBoxes
    .slice(0, n1)
    .map(([x2, y2, z2], n2) => ({
      n1,
      n2,
      d: Math.sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
    }))
  )
  .sort(({ d: d1 }, { d: d2 }) => d1 - d2);

const circuits = junctionBoxes.map((_, i) => new Set([i]));

distances.forEach(({ n1, n2 }, i) => {
  if (i === 1000) {
    console.log([...new Set([...Object.values(circuits)])].map(({ size }) => size).sort((a, b) => b - a).slice(0, 3).reduce((a, b) => a * b, 1));
  }
  const c1 = circuits[n1];
  const c2 = circuits[n2];
  if (c1 !== c2) {
    if (c1.size + c2.size === junctionBoxes.length) {
      console.log(junctionBoxes[n1][0] *  junctionBoxes[n2][0]);
      process.exit();
    }
    for (const n of c1) {
      circuits[n] = c2;
      c2.add(n);
    }
  }
});
