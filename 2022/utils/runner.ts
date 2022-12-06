import chalk from "chalk";

interface Options {
  title: string;
  day: number;
}

type Answer = number | string;
type SolutionCallbackSync = () => Answer;
type SolutionCallbackAsync = () => Promise<Answer>;

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
  const day = `Day ${opts.day}`;
  const header = `🎄 ${chalk.bold(day)} - ${chalk.bold.magenta(opts.title)} 🎄`;
  const remainingCharLength = 6 + 3; // trees + dash
  console.log(header);
  console.log(new Array(day.length + opts.title.length + remainingCharLength).fill("-").join(""));
  console.log();
};

const logPuzzle = (part: number, result: Answer, puzzle: PuzzlePartBase) => {
  console.log(`- Part ${part} -`);
  console.log(`${puzzle.message}:`, chalk.green(result));
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
