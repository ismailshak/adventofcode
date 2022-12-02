import fs from "node:fs";
import path from "node:path";
import { INPUT_FILE_NAME } from "@utils/constants";

export const getInputPath = (cwd: string) => {
  return path.join(cwd, INPUT_FILE_NAME);
};

export const parseInput = (cwd: string) =>
  fs.readFileSync(getInputPath(cwd), {
    encoding: "utf8",
  });
