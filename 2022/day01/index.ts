import { parseInput } from "@utils/input";

const input = parseInput(__dirname);

const isNewElf = (line: string) => line === "";
const hasMoreCalories = (currentTotal: number, largest: number) =>
  currentTotal > largest;
const getHigherCalories = (currentTotal: number, largest: number) =>
  hasMoreCalories(currentTotal, largest) ? currentTotal : largest;
const pushAndSort = (item: number, array: Array<number>) => {
  array.push(item);
  return array.sort((a, b) => b - a);
};

const part1 = () => {
  const mostCalories = input.split("\n").reduce(
    ({ largest, currentTotal }, line) => {
      if (isNewElf(line)) {
        return {
          largest: getHigherCalories(currentTotal, largest),
          currentTotal: 0,
        };
      }

      return { largest, currentTotal: currentTotal + parseInt(line, 10) };
    },
    { largest: 0, currentTotal: 0 }
  );

  console.log("Part 1:");
  console.log("Elf with the most calories has:", mostCalories.largest);
};

const part2 = () => {
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

  console.log("Part 2:");
  console.log(
    "Top three elves total",
    sortedTotals[0] + sortedTotals[1] + sortedTotals[2]
  );
};

part1();
console.log();
part2();
