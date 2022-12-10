import chalk from "chalk";
import {parseInput} from "@utils/input";
import {performance} from "perf_hooks";

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

const timeTakenToString = (start: number, end: number) => {
  const diff = end - start;
  return diff.toFixed(3);
};

const logPuzzle = async (part: number, input: string, puzzle: PuzzlePartSync | PuzzlePartAsync) => {
  console.log(`- Part ${part} -`);

  const start = performance.now();
  const result = await puzzle.solution(input);
  const end = performance.now();

  console.log(`${puzzle.message}:`, chalk.green(result), chalk.yellow(`(${timeTakenToString(start, end)}ms)`));
  //console.log(chalk.yellow(`(${timeTakenToString(start, end)}ms)`))
  console.log();
};

export const run = async (opts: Options, part1: PuzzlePartSync, part2: PuzzlePartSync) => {
  const input = parseInput(opts.cwd, opts.mock);
  logTitle(opts);

  await logPuzzle(1, input, part1);
  await logPuzzle(2, input, part2);
};

export const runAsync = async (opts: Options, part1: PuzzlePartAsync, part2: PuzzlePartAsync) => {
  const input = parseInput(opts.cwd, opts.mock);
  logTitle(opts);

  await logPuzzle(1, input, part1);
  await logPuzzle(2, input, part2);
};
