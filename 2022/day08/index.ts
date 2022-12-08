import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

interface Args {
  lines: Array<string>;
  currentLine: string;
  lineIndex: number;
  charIndex: number;
}

const isTreeVisible = (currentHeight: number, treeHeight: string) => currentHeight <= parseInt(treeHeight, 10);

const isTopVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = lineIndex - 1; i >= 0; i--) {
    if (isTreeVisible(currentHeight, lines[i][charIndex])) {
      return false;
    }
  }

  return true;
};

const isBottomVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = lineIndex + 1; i < lines.length; i++) {
    if (isTreeVisible(currentHeight, lines[i][charIndex])) {
      return false;
    }
  }

  return true;
};

const isLeftVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = charIndex - 1; i >= 0; i--) {
    if (isTreeVisible(currentHeight, lines[lineIndex][i])) {
      return false;
    }
  }

  return true;
};

const isRightVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = charIndex + 1; i < currentLine.length; i++) {
    if (isTreeVisible(currentHeight, lines[lineIndex][i])) {
      return false;
    }
  }

  return true;
};

const isVisible = (args: Args) =>
  isTopVisible(args) || isBottomVisible(args) || isLeftVisible(args) || isRightVisible(args);

const countVisibleInRow = (lines: Array<string>, currentLine: string, lineIndex: number) => {
  let count = 0;

  // Start at index 1, and stop at the second to last character
  // (so that we can skip the edges)
  for (let i = 1; i < currentLine.length - 1; i++) {
    if (isVisible({lines, currentLine, lineIndex, charIndex: i})) {
      count++;
    }
  }

  return count;
};

// Count trees on the perimeter of the grid (subtracting 4 because of the 4 overlapping corner trees)
const countPerimeterTrees = (lines: Array<string>) => lines.length * 2 + lines[0].length * 2 - 4;

const part1 = (input: string) => {
  const lines = inputToLines(input);

  const defaultVisible = countPerimeterTrees(lines);

  console.log(defaultVisible);

  const innerVisibleCount = lines.reduce((acc, line, index) => {
    if (index === 0 || index === lines.length - 1) return acc;

    return acc + countVisibleInRow(lines, line, index);
  }, 0);

  return defaultVisible + innerVisibleCount;
};

const part2 = (input: string) => 0;

run(
  {cwd: __dirname, day: 8, title: "Treetop Tree House"},
  {solution: part1, message: "Number of trees visible from the outside grid"},
  {solution: part2, message: "placeholder"}
);
