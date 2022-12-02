import { parseInput } from "@utils/input";
import { mod } from "@utils/math";

const input = parseInput(__dirname);

enum MoveScore {
  ROCK = 1,
  PAPER = 2,
  SCISSORS = 3,
}

const inputToMove: Record<string, MoveScore> = {
  X: MoveScore.ROCK,
  A: MoveScore.ROCK,
  Y: MoveScore.PAPER,
  B: MoveScore.PAPER,
  Z: MoveScore.SCISSORS,
  C: MoveScore.SCISSORS,
};

enum RoundResult {
  WIN = 6,
  DRAW = 3,
  LOSS = 0,
}

// For a response action to win against the opponent action, it's always going to be
// the previous action in the enumerated version of the moves above
// i.e. if response is PAPER, opponent has to be the previous move which is ROCK
// i.e. if response is ROCK, opponent has to be the previous move which is SCISSOR (cycling around)
const isWin = (opponent: MoveScore, response: MoveScore) =>
  mod(response - opponent, 3) === 1;

const getGameResult = (
  opponent: MoveScore,
  response: MoveScore
): RoundResult => {
  if (opponent === response) {
    return RoundResult.DRAW;
  }

  if (isWin(opponent, response)) {
    return RoundResult.WIN;
  }

  return RoundResult.LOSS;
};

const inputToScore = (value: string) => {
  if (!(value in inputToMove)) {
    throw `Input parse error. Move value '${value}' not in map`;
  }

  return inputToMove[value];
};

const roundToMoves = (round: string) => {
  const parts = round.split(" ");
  return { opponent: inputToScore(parts[0]), response: inputToScore(parts[1]) };
};

const part1 = () => {
  const rounds = input.trim().split("\n");

  const totalScore = rounds.reduce((acc, round) => {
    const { opponent, response } = roundToMoves(round);

    const gameResult = getGameResult(opponent, response);
    const roundResult = gameResult + response;
    const totalResult = acc + roundResult;

    return totalResult;
  }, 0);

  console.log("Part 1: Total score is", totalScore);
};

part1();
