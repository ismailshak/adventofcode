import {run} from "@utils/runner";

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

const part1 = (input: string) => {
  const sacks = input.trim().split("\n");

  const sum = sacks.reduce((acc, sack) => {
    const {first, second} = getCompartments(sack);
    const commonItem = findCommonItem(first, second);
    const priority = getPriority(commonItem);

    return acc + priority;
  }, 0);

  return sum;
};

const findBadge = (sack1: string, sack2: string, sack3: string) => {
  const setA = new Set(sack1.split(""));
  const setB = new Set(sack2.split(""));
  const setC = new Set(sack3.split(""));
  const intersection = new Set();

  for (const item of setA) {
    if (setB.has(item) && setC.has(item)) {
      intersection.add(item);
      break; // Assuming there's only ever 1 common item
    }
  }

  if (intersection.size !== 1) {
    throw `Intersection set did not have exactly 1 item: ${Array.from(intersection)}`;
  }

  return intersection.values().next().value as string;
};

const part2 = (input: string) => {
  const sacks = input.trim().split("\n");

  if (sacks.length % 3 !== 0) throw `Input is not divisible by 3`;

  const {sum} = sacks.reduce(
    (acc, sack, currentIndex) => {
      if (acc.elfCount === 2) {
        const badge = findBadge(sacks[currentIndex - 2], sacks[currentIndex - 1], sack);
        const priority = getPriority(badge);
        return {
          sum: acc.sum + priority,
          elfCount: 0,
        };
      }

      return {sum: acc.sum, elfCount: acc.elfCount + 1};
    },
    {sum: 0, elfCount: 0}
  );

  return sum;
};

run(
  {cwd: __dirname, day: 3, title: "Rucksack Reorganization"},
  {solution: part1, message: "Sum of common items"},
  {solution: part2, message: "Sum of badge priorities"}
);
