import {inputToLines, parseInput} from "@utils/input";
import {run} from "@utils/runner";

type SectionRange = [number, number];

const input = parseInput(__dirname);

const parseAssignment = (value: string) => value.split("-").map((id) => parseInt(id, 10)) as SectionRange;

const getAssignments = (pair: string) => {
  const assignments = pair.split(",");
  const first = parseAssignment(assignments[0]);
  const second = parseAssignment(assignments[1]);

  return {first, second};
};

// Whether range B is fully contained within range A (or vice versa)
const isInRange = (rangeA: SectionRange, rangeB: SectionRange) =>
  (rangeA[0] <= rangeB[0] && rangeA[1] >= rangeB[1]) || (rangeA[0] >= rangeB[0] && rangeA[1] <= rangeB[1]);

const part1 = () => {
  const pairs = inputToLines(input);

  const total = pairs.reduce((acc, pair) => {
    const {first, second} = getAssignments(pair);
    const insideRange = isInRange(first, second);

    if (insideRange) {
      return acc + 1;
    }

    return acc;
  }, 0);

  return total;
};

const part2 = () => 0;

run(
  {day: 2, title: "Camp Cleanup"},
  {solution: part1, message: "Number of pairs inside range"},
  {solution: part2, message: "placeholder"}
);
