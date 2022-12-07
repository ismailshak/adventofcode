import chalk from "chalk";
import {parseInput} from "@utils/input";

interface Options {
  cwd: string;
  day: number;
  title: string;
  mock?: boolean;
}

type Answer = number | string;
type SolutionCallbackSync = (input: string) => Answer;
type SolutionCallbackAsync = (input: string) => Promise<Answer>;

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
  const input = parseInput(opts.cwd, opts.mock);
  logTitle(opts);

  logPuzzle(1, part1.solution(input), part1);
  logPuzzle(2, part2.solution(input), part2);
};

export const runAsync = async (opts: Options, part1: PuzzlePartAsync, part2: PuzzlePartAsync) => {
  const input = parseInput(opts.cwd, opts.mock);
  logTitle(opts);

  logPuzzle(1, await part1.solution(input), part1);
  logPuzzle(2, await part2.solution(input), part2);
};
