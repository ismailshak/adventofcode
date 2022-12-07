import {run} from "@utils/runner";
import Stack from "@utils/stack";

type CrateStack = Stack<string>;

interface Column {
  /* the position of the number character in the column row */
  index: number;
  /* the actual number in the column row */
  number: number;
}

interface Move {
  count: number;
  from: number;
  to: number;
}

const getColumns = (row: string) => {
  const columns: Array<Column> = [];
  for (let i = 0; i < row.length; i++) {
    const parsedChar = parseInt(row[i], 10);
    if (parsedChar) {
      columns.push({index: i, number: parsedChar});
    }
  }

  return columns;
};

const generateStacks = (count: number) => {
  const stacks: Array<CrateStack> = [];

  for (let i = 0; i < count; i++) {
    stacks.push(new Stack<string>());
  }

  return stacks;
};

// Coerces empty strings or whitespaces into undefined.
// Assumption: the column number index is aligned with the letter in the crate letter
// i.e.
// [D]
//  1
const parseCrateRow = (columns: Array<Column>, row: string) =>
  columns.map((column) => (row[column.index]?.trim() ? row[column.index] : undefined));

const parseCrateSetup = (setup: Array<string>) => {
  const crateSection = setup.slice(0, setup.length - 1);
  const columnLine = setup[setup.length - 1];

  const columns = getColumns(columnLine);

  const stacks = generateStacks(columns.length);

  crateSection.reverse().forEach((row) => {
    const crates = parseCrateRow(columns, row);
    crates.forEach((crate, index) => {
      if (!crate) {
        return;
      }

      stacks[index].push(crate);
    });
  });

  return stacks;
};

const chunkInput = (input: string) => {
  const lines = input.split("\n");

  let splitIndex = 0;

  for (let i = 0; i < lines.length; i++) {
    if (lines[i].includes("move")) {
      splitIndex = i - 1;

      // subtracting 1 will land on the empty line between the two sections
      // so this is a sanity check
      if (lines[splitIndex] !== "") {
        throw `Split index wasn't actually the empty line: ${lines[splitIndex]}`;
      }

      break;
    }
  }

  return {
    setup: lines.slice(0, splitIndex),
    moves: lines.slice(splitIndex + 1), // moves will be the line after the empty separator line
  };
};

const parseMove = (move: string): Move | undefined => {
  const nums = move.replaceAll(/[^0-9]/g, ",");

  if (!nums) {
    return;
  }

  const parts = nums.split(",").filter((num) => !!num);
  return {
    count: parseInt(parts[0], 10),
    from: parseInt(parts[1], 10) - 1, // index starts at 0, but columns start at 1
    to: parseInt(parts[2], 10) - 1, // index starts at 0, but columns start at 1
  };
};

const part1 = (input: string) => {
  const {setup, moves} = chunkInput(input);

  const stacks = parseCrateSetup(setup);

  moves.forEach((rawMove) => {
    const move = parseMove(rawMove);
    if (!move) return;

    for (let i = 0; i < move.count; i++) {
      stacks[move.to].push(stacks[move.from].pop());
    }
  });

  return stacks.map((stack) => stack.peek()).join("");
};

const part2 = (input: string) => {
  const {setup, moves} = chunkInput(input);

  const stacks = parseCrateSetup(setup);

  moves.forEach((rawMove) => {
    const move = parseMove(rawMove);
    if (!move) return;

    const crateHolder: Array<string> = [];
    for (let i = 0; i < move.count; i++) {
      crateHolder.push(stacks[move.from].pop());
    }

    stacks[move.to].pushMany(crateHolder.reverse());
  });

  return stacks.map((stack) => stack.peek()).join("");
};

run(
  {cwd: __dirname, day: 5, title: "Supply Stacks"},
  {solution: part1, message: "Crates on the top of each stack"},
  {solution: part2, message: "Crates on the top of each stack with updated rules"}
);
