import {mod} from "@utils/math";
import {run} from "@utils/runner";

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

const inputToResult: Record<string, RoundResult> = {
  X: RoundResult.LOSS,
  Y: RoundResult.DRAW,
  Z: RoundResult.WIN,
};

// For a response action to win against the opponent action, it's always going to be
// the previous action in the enumerated version of the moves above
// i.e. if response is PAPER, opponent has to be the previous move which is ROCK
// i.e. if response is ROCK, opponent has to be the previous move which is SCISSOR (cycling around)
const isWin = (opponent: MoveScore, response: MoveScore) => mod(response - opponent, 3) === 1;

const getGameResult = (opponent: MoveScore, response: MoveScore): RoundResult => {
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
  return {opponent: inputToScore(parts[0]), response: inputToScore(parts[1])};
};

const part1 = (input: string) => {
  const rounds = input.trim().split("\n");

  const totalScore = rounds.reduce((acc, round) => {
    const {opponent, response} = roundToMoves(round);

    const gameResult = getGameResult(opponent, response);
    const roundResult = gameResult + response;
    const totalResult = acc + roundResult;

    return totalResult;
  }, 0);

  return totalScore;
};

const getInputToResult = (value: string) => {
  if (!(value in inputToResult)) {
    throw `Input parse error. Result value '${value}' not in map`;
  }

  return inputToResult[value];
};

const roundToAction = (round: string) => {
  const parts = round.split(" ");
  return {
    opponent: inputToScore(parts[0]),
    result: getInputToResult(parts[1]),
  };
};

// To get the losing move, we need to find the previous action in the enumerated sequence
// we need to start at 1 (so we subtract 2), then we add 1 back so that we just subtract 1 overall
const getLosingMove = (opponent: MoveScore) => mod(opponent - 2, 3) + 1;
// To get the winning move, we need the next action in the enumerated sequence, so we add 1 to the mod result
const getWinningMove = (opponent: MoveScore) => mod(opponent, 3) + 1;

const getResponseMove = (opponent: MoveScore, result: RoundResult) => {
  if (result === RoundResult.DRAW) {
    return opponent;
  }

  if (result === RoundResult.LOSS) {
    return getLosingMove(opponent);
  }

  return getWinningMove(opponent);
};

const part2 = (input: string) => {
  const rounds = input.trim().split("\n");

  const totalScore = rounds.reduce((acc, round) => {
    const {opponent, result} = roundToAction(round);

    const response = getResponseMove(opponent, result);
    const roundScore = response + result;
    const totalResult = acc + roundScore;

    return totalResult;
  }, 0);

  return totalScore;
};

run(
  {cwd: __dirname, day: 2, title: "Rock Paper Scissors"},
  {solution: part1, message: "Total score"},
  {solution: part2, message: "Total score"}
);
