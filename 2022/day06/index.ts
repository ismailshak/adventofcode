import {inputToLines, parseInput} from "@utils/input";
import {run} from "@utils/runner";

const input = parseInput(__dirname);

const arePreviousFourUnique = (index: number) => {
  const substr = input.slice(index - 4, index);
  return new Set(substr.split("")).size === 4;
};

const part1 = () => {
  for (let i = 0; i < input.length; i++) {
    if (i < 4) {
      continue;
    }

    if (arePreviousFourUnique(i)) {
      return i;
    }
  }

  return 0;
};

const part2 = () => 0;

run(
  {day: 6, title: "Tuning Trouble"},
  {solution: part1, message: "placeholder"},
  {solution: part2, message: "placeholder"}
);
