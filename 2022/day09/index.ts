import {inputToLines} from "@utils/input";
import {run} from "@utils/runner";

type Coords = {x: number; y: number};

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

const moveHead = (direction: Direction, head: Coords) => {
  switch (direction) {
    case Direction.UP:
      head.y++;
      break;
    case Direction.DOWN:
      head.y--;
      break;
    case Direction.LEFT:
      head.x--;
      break;
    case Direction.RIGHT:
      head.x++;
      break;
    default:
      throw `${direction} did not find a matching switch case for head`;
  }

  return head;
};

const moveTail = (head: Coords, tail: Coords) => {
  const diffX = Math.abs(head.x - tail.x);
  const diffY = Math.abs(head.y - tail.y);

  if (diffX < 2 && diffY < 2) {
    return {current: tail, parent: head};
  }

  if (diffX > 1 && !diffY) {
    tail.x += head.x - tail.x > 0 ? 1 : -1;
  } else if (diffY > 1 && !diffX) {
    tail.y += head.y - tail.y > 0 ? 1 : -1;
  } else {
    tail.x += head.x - tail.x > 0 ? 1 : -1;
    tail.y += head.y - tail.y > 0 ? 1 : -1;
  }

  return {current: tail, parent: head};
};

const moveKnot = (direction: Direction, parent: Coords, current: Coords) => {
  const head = moveHead(direction, parent);
  return moveTail(head, current);
};

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
    const {current, parent} = moveKnot(direction, {x: this.hX, y: this.hY}, {x: this.tX, y: this.tY});

    this.hX = parent.x;
    this.hY = parent.y;
    this.tX = current.x;
    this.tY = current.y;

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

const moveRope = (move: Move, rope: Array<Coords>, set: Set<string>) => {
  for (let i = 0; i < move.count; i++) {
    rope.forEach((knot, index) => {
      if (index === 0) {
        const head = moveHead(move.direction, knot);
        rope[0] = head;
        return;
      }

      const {current} = moveTail(rope[index - 1], knot);
      if (index === rope.length - 1) {
        set.add(`${current.x},${current.y}`);
      }
    });
  }
};

const part2 = (input: string) => {
  const ropeSize = 10;
  const lines = inputToLines(input);
  const tailPositions = new Set<string>();
  const knots: Array<Coords> = [];
  for (let i = 0; i < ropeSize; i++) {
    knots[i] = {x: 11, y: 5};
  }

  lines.forEach((line) => {
    const move = parseMove(line);
    moveRope(move, knots, tailPositions);
  });
  return tailPositions.size;
};

run(
  {cwd: __dirname, day: 9, title: "Rope Bridge", mock: false},
  {solution: part1, message: "Number of positions tail of rope visited"},
  {solution: part2, message: "placeholder"}
);
