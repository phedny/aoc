export function getInput<S extends string>(file: string, spec: S): Spec<S>;
export function getInput<S extends string[]>(file: string, ...specs: S): {[I in keyof S]: Spec<S[I]>};

type Spec<S extends string> = S extends ''
  ? string
  : S extends 'I'
  ? number
  : S extends `${'C'|'L'|'B'|`S${string}`}${infer R}`
  ? Spec<R>[]
  : S extends `=${string}`
  ? boolean
  : never;
