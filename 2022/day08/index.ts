import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

interface Args {
  lines: Array<string>;
  currentLine: string;
  lineIndex: number;
  charIndex: number;
}

const isTreeBlocked = (currentHeight: number, treeHeight: string) => currentHeight <= parseInt(treeHeight, 10);

const isTopVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = lineIndex - 1; i >= 0; i--) {
    if (isTreeBlocked(currentHeight, lines[i][charIndex])) {
      return false;
    }
  }

  return true;
};

const isBottomVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = lineIndex + 1; i < lines.length; i++) {
    if (isTreeBlocked(currentHeight, lines[i][charIndex])) {
      return false;
    }
  }

  return true;
};

const isLeftVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = charIndex - 1; i >= 0; i--) {
    if (isTreeBlocked(currentHeight, lines[lineIndex][i])) {
      return false;
    }
  }

  return true;
};

const isRightVisible = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  const currentHeight = parseInt(currentLine[charIndex], 10);
  for (let i = charIndex + 1; i < currentLine.length; i++) {
    if (isTreeBlocked(currentHeight, lines[lineIndex][i])) {
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

  const innerVisibleCount = lines.reduce((acc, line, index) => {
    if (index === 0 || index === lines.length - 1) return acc;

    return acc + countVisibleInRow(lines, line, index);
  }, 0);

  return defaultVisible + innerVisibleCount;
};

const topScore = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  if (lineIndex === 0) return 0;

  const currentHeight = parseInt(currentLine[charIndex], 10);
  const firstIndex = lineIndex - 1;
  for (let i = firstIndex; i >= 0; i--) {
    if (isTreeBlocked(currentHeight, lines[i][charIndex])) {
      return firstIndex - i + 1;
    }
  }

  return lineIndex;
};

const bottomScore = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  if (lineIndex === lines.length - 1) return 0;

  const currentHeight = parseInt(currentLine[charIndex], 10);
  const firstIndex = lineIndex + 1;
  for (let i = firstIndex; i < lines.length; i++) {
    if (isTreeBlocked(currentHeight, lines[i][charIndex])) {
      return i - firstIndex + 1;
    }
  }

  return lines.length - 1 - lineIndex;
};

const leftScore = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  if (charIndex === 0) return 0;

  const currentHeight = parseInt(currentLine[charIndex], 10);
  const firstIndex = charIndex - 1;
  for (let i = firstIndex; i >= 0; i--) {
    if (isTreeBlocked(currentHeight, lines[lineIndex][i])) {
      return firstIndex - i + 1;
    }
  }

  return charIndex;
};

const rightScore = ({lines, currentLine, lineIndex, charIndex}: Args) => {
  if (charIndex === currentLine.length - 1) return 0;

  const currentHeight = parseInt(currentLine[charIndex], 10);
  const firstIndex = charIndex + 1;
  for (let i = charIndex + 1; i < currentLine.length; i++) {
    if (isTreeBlocked(currentHeight, lines[lineIndex][i])) {
      return i - firstIndex + 1;
    }
  }

  return currentLine.length - 1 - charIndex;
};

const calculateScenicScore = (args: Args) => topScore(args) * bottomScore(args) * leftScore(args) * rightScore(args);

const findBestScoreInRow = (lines: Array<string>, currentLine: string, lineIndex: number) => {
  let highestScore = 0;

  for (let i = 1; i < currentLine.length; i++) {
    const currentTreeScore = calculateScenicScore({lines, currentLine, lineIndex, charIndex: i});
    if (currentTreeScore > highestScore) {
      highestScore = currentTreeScore;
    }
  }

  return highestScore;
};

const part2 = (input: string) => {
  const lines = inputToLines(input);

  const bestScenicScore = lines.reduce((acc, line, index) => {
    const bestScore = findBestScoreInRow(lines, line, index);
    if (bestScore > acc) {
      return bestScore;
    }

    return acc;
  }, 0);

  return bestScenicScore;
};

run(
  {cwd: __dirname, day: 8, title: "Treetop Tree House", mock: false},
  {solution: part1, message: "Number of trees visible from outside the grid"},
  {solution: part2, message: "Highest scenic score"}
);
