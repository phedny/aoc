import { readFileSync } from 'fs';

export function getInput(file, ...specs) {
  const data = readFileSync(file, { encoding: 'utf-8' });
  if (specs.length === 1) {
    return parseBlock(data, specs[0]);
  } else {
    return data.split('\n\n').map((data, i) => parseBlock(data, specs[i]));
  }
}

function parseBlock(data, spec) {
  if (spec.length === 0) {
    return data;
  }
  switch (spec[0]) {
    case 'C':
      return data.split('').map((c) => parseBlock(c, spec.slice(1)));
    case 'I':
      return parseInt(data);
    case 'L':
      return data.split('\n').map((line) => parseBlock(line, spec.slice(1)));
    case 'B':
      return data.split('\n\n').map((line) => parseBlock(line, spec.slice(1)));
    case 'S':
      return data.split(spec[1]).map((part) => parseBlock(part, spec.slice(2)));
    case '=':
      return data === spec.slice(1);
  }
}
