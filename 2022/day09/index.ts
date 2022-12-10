import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

enum Direction {
  UP = "U",
  DOWN = "D",
  LEFT = "L",
  RIGHT = "R",
}

interface Move {
  direction: Direction;
  count: number;
}

class FakeGrid {
  private set: Set<string>;
  private hX: number;
  private hY: number;
  private tX: number;
  private tY: number;

  constructor() {
    this.set = new Set();
    (this.hX = 0), (this.hY = 0);
    (this.tX = 0), (this.tY = 0);
    this.markVisited();
  }

  moveHead(move: Move) {
    for (let i = 0; i < move.count; i++) {
      this.move(move.direction);
    }
  }

  move(direction: Direction) {
    switch (direction) {
      case Direction.UP:
        this.hY++;
        if (Math.abs(this.tY - this.hY) > 1) {
          this.tX = this.hX;
          this.tY = this.hY - 1;
        }
        break;
      case Direction.DOWN:
        this.hY--;
        if (Math.abs(this.tY - this.hY) > 1) {
          this.tX = this.hX;
          this.tY = this.hY + 1;
        }
        break;
      case Direction.LEFT:
        this.hX--;
        if (Math.abs(this.tX - this.hX) > 1) {
          this.tY = this.hY;
          this.tX = this.hX + 1;
        }
        break;
      case Direction.RIGHT:
        this.hX++;
        if (Math.abs(this.tX - this.hX) > 1) {
          this.tY = this.hY;
          this.tX = this.hX - 1;
        }
        break;
      default:
        throw `${direction} did not find a matching switch case`;
    }

    this.markVisited();
  }

  markVisited(x = this.tX, y = this.tY) {
    this.set.add(`${x},${y}`);
  }

  getVisitedCount() {
    return this.set.size;
  }
}

const parseMove = (line: string): Move => {
  const parts = line.split(" ");
  const count = parseInt(parts[1], 10);

  switch (parts[0]) {
    case "U":
      return {direction: Direction.UP, count};
    case "D":
      return {direction: Direction.DOWN, count};
    case "L":
      return {direction: Direction.LEFT, count};
    case "R":
      return {direction: Direction.RIGHT, count};
  }

  throw `${parts[0]} did not match switch when parsing`;
};

const part1 = (input: string) => {
  const lines = inputToLines(input);
  const fakeGrid = new FakeGrid();

  lines.forEach((line) => {
    const move = parseMove(line);
    fakeGrid.moveHead(move);
  });

  return fakeGrid.getVisitedCount();
};

const part2 = (input: string) => 0;

run(
  {cwd: __dirname, day: 9, title: "Rope Bridge", mock: false},
  {solution: part1, message: "Number of positions tail of rope visited"},
  {solution: part2, message: "placeholder"}
);
