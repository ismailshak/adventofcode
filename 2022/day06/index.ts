import {parseInput} from "@utils/input";
import {run} from "@utils/runner";

const input = parseInput(__dirname);

const arePreviousXUnique = (index: number, count: number) => {
  const substr = input.slice(index - count, index);
  return new Set(substr.split("")).size === count;
};

const part1 = () => {
  const uniqueCount = 4;
  for (let i = 0; i < input.length; i++) {
    if (i < uniqueCount) {
      continue;
    }

    if (arePreviousXUnique(i, uniqueCount)) {
      return i;
    }
  }

  return 0;
};

const part2 = () => {
  const uniqueCount = 14;
  for (let i = 0; i < input.length; i++) {
    if (i < uniqueCount) {
      continue;
    }

    if (arePreviousXUnique(i, uniqueCount)) {
      return i;
    }
  }

  return 0;
};

run(
  {day: 6, title: "Tuning Trouble"},
  {solution: part1, message: "Characters processed before 4 unique found"},
  {solution: part2, message: "Characters processed before 14 unique found"}
);
