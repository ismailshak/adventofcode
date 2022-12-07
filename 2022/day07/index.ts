import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

const DIR_PLACEHOLDER = -1;

interface File {
  name: string;
  size: number;
}

interface Directory {
  name: string;
}

interface Command {
  command: string;
  arg: string;
}

const isCommand = (line: string) => line.startsWith("$");
const isCd = (line: string) => line.startsWith("$ cd");
const isDir = (line: string) => line.startsWith("dir");

const parseCommand = (line: string): Command => {
  const parts = line.split(" ");
  return {command: parts[1], arg: parts[2]};
};

const parseDir = (line: string): Directory => {
  const parts = line.split(" ");
  return {name: parts[1]};
};

const parseFile = (line: string): File => {
  const parts = line.split(" ");
  return {size: parseInt(parts[0], 10), name: parts[1]};
};

// Builds an object with each of its properties set to a path. Each one of those paths will include all the files,
// and any directories will be marked with a placeholder (-1), so that we can account for it when summing up totals.
const buildTree = (lines: Array<string>) => {
  const fs: any = {};
  const pwd: Array<string> = [];

  lines.forEach((line) => {
    if (isCommand(line) && isCd(line)) {
      const {arg} = parseCommand(line);
      if (arg === "..") {
        pwd.pop();
        return;
      }

      pwd.push(arg);
      return;
    }

    if (!isCommand(line)) {
      const path = pwd.join("/");
      if (!(path in fs)) {
        fs[path] = {};
      }

      if (isDir(line)) {
        const {name} = parseDir(line);
        fs[path][name] = DIR_PLACEHOLDER;
        return;
      }

      const {name, size} = parseFile(line);
      fs[path][name] = size;
    }
  });

  return fs;
};

// For every path in the tree, sum up the totals. If a directory is found inside a path,
// we sum that directory and lift it's value to the top
const sumTree = (tree: Record<string, any>, totals: Record<string, any>, dir: string): number => {
  const total: number = Object.entries<number>(tree[dir]).reduce((acc, [file, size]) => {
    if (size === DIR_PLACEHOLDER) {
      return acc + sumTree(tree, totals, `${dir}/${file}`);
    }

    return acc + size;
  }, 0);

  totals[dir] = total;
  return total;
};

const part1 = (input: string) => {
  const lines = inputToLines(input);

  const fs = buildTree(lines);

  const totals = {};
  sumTree(fs, totals, "/");

  // Filter out all values that are greater than threshold then sum them
  return Object.values<number>(totals)
    .filter((num) => num <= 100000)
    .reduce((sum, num) => sum + num, 0);
};

const part2 = (input: string) => {
  const lines = inputToLines(input);

  const fs = buildTree(lines);

  const totals: any = {};
  sumTree(fs, totals, "/");

  const spaceNeeded = 30000000 - (70000000 - totals["/"]);

  return Object.values<number>(totals)
    .filter((num) => num >= spaceNeeded)
    .sort((a, b) => (a > b ? 1 : -1))[0];
};

run(
  {cwd: __dirname, day: 7, title: "No Space Left On Device"},
  {solution: part1, message: "Total size of directories with less than 100000 space"},
  {solution: part2, message: "placeholder"}
);
