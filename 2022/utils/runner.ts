interface Options {
  title: string;
  day: number;
}

type SolutionCallbackSync = () => number;
type SolutionCallbackAsync = () => Promise<number>;

interface PuzzlePartBase {
  message: string;
}

interface PuzzlePartSync extends PuzzlePartBase {
  solution: SolutionCallbackSync;
}

interface PuzzlePartAsync extends PuzzlePartBase {
  solution: SolutionCallbackAsync;
}

const logTitle = (opts: Options) => {
  const header = `Day ${opts.day} - ${opts.title}`;
  console.log(header);
  console.log(new Array(header.length).fill("-").join(""));
  console.log();
};

const logPuzzle = (part: number, result: number, puzzle: PuzzlePartBase) => {
  console.log(`- Part ${part} -`);
  console.log(`${puzzle.message}:`, result);
  console.log();
};

export const run = (opts: Options, part1: PuzzlePartSync, part2: PuzzlePartSync) => {
  logTitle(opts);

  logPuzzle(1, part1.solution(), part1);
  logPuzzle(2, part2.solution(), part2);
};

export const runAsync = async (opts: Options, part1: PuzzlePartAsync, part2: PuzzlePartAsync) => {
  logTitle(opts);

  logPuzzle(1, await part1.solution(), part1);
  logPuzzle(2, await part2.solution(), part2);
};
