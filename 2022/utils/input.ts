import fs from "node:fs";
import path from "node:path";
import {INPUT_FILE_NAME, TEST_INPUT_FILE_NAME} from "@utils/constants";

export const getInputPath = (cwd: string, isTest = false) => {
  return path.join(cwd, isTest ? TEST_INPUT_FILE_NAME : INPUT_FILE_NAME);
};

export const parseInput = (cwd: string, isTest = false) =>
  fs.readFileSync(getInputPath(cwd, isTest), {
    encoding: "utf8",
  });
