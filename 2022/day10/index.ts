import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

enum Command {
  noop = "noop",
  add = "addx",
}

const parseCommand = (line: string) => {
  const parts = line.split(" ");
  if (parts[0] === "noop") {
    return {command: Command.noop};
  }

  return {command: Command.add, value: parseInt(parts[1], 10)};
};

const handleCommand = (cycles: Array<number>, command: Command, value = 0) => {
  const previousValue = cycles.at(-1) as number;
  if (command === Command.noop) {
    cycles.push(previousValue);
    return;
  }

  cycles.push(previousValue);
  cycles.push(previousValue + value);
};

const getStrengths = (cycles: Array<number>) => {
  const strengths: Array<number> = [];

  strengths.push(cycles[19] * 20);
  strengths.push(cycles[59] * 60);
  strengths.push(cycles[99] * 100);
  strengths.push(cycles[139] * 140);
  strengths.push(cycles[179] * 180);
  strengths.push(cycles[219] * 220);

  return strengths;
};

const part1 = (input: string) => {
  const commands = inputToLines(input);

  const registerPerCycle: Array<number> = [1];

  commands.forEach((line) => {
    const {command, value} = parseCommand(line);
    handleCommand(registerPerCycle, command, value);
  });

  const strengths = getStrengths(registerPerCycle);

  return strengths.reduce((acc, strength) => acc + strength, 0);
};

const generateSprites = (cycles: Array<number>) =>
  cycles.map((_: number, i: number) => {
    return new Array(3).fill(cycles[i] - 1)
      .map((x, i) => x + i)
      .includes(i % 40) ? '#' : '.'
  })

const drawSprites = (sprites: Array<"#" | ".">) =>
  sprites.reduce((acc, char, index) => acc + char + ((index + 1) % 40 === 0 ? "\n" : ""), "")

const part2 = (input: string) => {
  const commands = inputToLines(input);
  const registerPerCycle: Array<number> = [1];

  commands.forEach((line) => {
    const {command, value} = parseCommand(line);
    handleCommand(registerPerCycle, command, value);
  });

  const sprites = generateSprites(registerPerCycle);
  const drawing = drawSprites(sprites);
  return "\n" + drawing + "\n";
};

run(
  {cwd: __dirname, day: 10, title: "Cathode-Ray Tube", mock: false},
  {solution: part1, message: "Sum of signal strengths"},
  {solution: part2, message: "Eight capital letters produced on CRT"}
);
