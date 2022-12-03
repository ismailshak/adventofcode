import {parseInput} from "@utils/input";
import {run} from "@utils/runner";

const input = parseInput(__dirname);

const getCompartments = (sack: string) => {
  const first = sack.substring(0, sack.length / 2);
  const second = sack.substring(sack.length / 2);

  return {first, second};
};

const findCommonItem = (first: string, second: string) => {
  const setA = new Set(first.split(""));
  const setB = new Set(second.split(""));
  const intersection = new Set();

  for (const item of setA) {
    if (setB.has(item)) {
      intersection.add(item);
      break; // Assuming there's only ever 1 common item
    }
  }

  if (intersection.size !== 1) {
    throw `Intersection set did not have exactly 1 item: ${Array.from(intersection)}`;
  }

  return intersection.values().next().value as string;
};

const getPriority = (letter: string) => {
  if (letter.length > 1) throw `Letter '${letter}' was not of length 1`;

  const code = letter.charCodeAt(0);

  // 'a' <--> 'z'
  if (code >= 97 && code <= 122) {
    return code - 96;
  }

  // 'A' <--> 'Z'
  if (code >= 65 && code <= 90) {
    return code - 38;
  }

  throw `Character code '${code}' was not within the expected range`;
};

const part1 = () => {
  const sacks = input.trim().split("\n");

  const sum = sacks.reduce((acc, sack) => {
    const {first, second} = getCompartments(sack);
    const commonItem = findCommonItem(first, second);
    const priority = getPriority(commonItem);

    return acc + priority;
  }, 0);

  return sum;
};

run(
  {day: 3, title: "Rucksack Reorganization"},
  {solution: part1, message: "Sum of common items"},
  {solution: () => 0, message: "placeholder"}
);
