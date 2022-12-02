import fs from "node:fs";
import path from "node:path";

export const getInputPath = (cwd: string) => {
  return path.join(cwd, "input.txt");
};

export const parseInput = (cwd: string) =>
  fs.readFileSync(getInputPath(cwd), {
    encoding: "utf8",
  });
