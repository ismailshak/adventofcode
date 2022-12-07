import {run} from "@utils/runner";

const isNewElf = (line: string) => line === "";
const hasMoreCalories = (currentTotal: number, largest: number) => currentTotal > largest;
const getHigherCalories = (currentTotal: number, largest: number) =>
  hasMoreCalories(currentTotal, largest) ? currentTotal : largest;
const pushAndSort = (item: number, array: Array<number>) => {
  array.push(item);
  return array.sort((a, b) => b - a);
};

const part1 = (input: string) => {
  const mostCalories = input.split("\n").reduce(
    ({largest, currentTotal}, line) => {
      if (isNewElf(line)) {
        return {
          largest: getHigherCalories(currentTotal, largest),
          currentTotal: 0,
        };
      }

      return {largest, currentTotal: currentTotal + parseInt(line, 10)};
    },
    {largest: 0, currentTotal: 0}
  );

  return mostCalories.largest;
};

const part2 = (input: string) => {
  let sortedTotals: Array<number> = [];
  let singleElfTotal = 0;
  const lines = input.split("\n");
  lines.forEach((line) => {
    if (isNewElf(line)) {
      sortedTotals = pushAndSort(singleElfTotal, sortedTotals);
      singleElfTotal = 0;
      return;
    }

    singleElfTotal += parseInt(line, 10);
  });

  return sortedTotals[0] + sortedTotals[1] + sortedTotals[2];
};

run(
  {cwd: __dirname, day: 1, title: "Calorie Counting"},
  {solution: part1, message: "Elf with the most calories has"},
  {solution: part2, message: "Sum of calories for top 3 elves"}
);
