import {execSync} from "node:child_process";
import fs, {existsSync} from "node:fs";
import path from "node:path";
import {digitPrompt, textPrompt} from "./prompts";

const getIndexContent = (day: number, title: string) => `import { inputToLines, parseInput } from "@utils/input";
import { run } from "@utils/runner";

const input = parseInput(__dirname, true);

const part1 = () => 0

const part2 = () => 0;

run(
  {day: ${day}, title: "${title}"},
  {solution: part1, message: "placeholder"},
  {solution: part2, message: "placeholder"}
);

`;

const getReadmeContent = (day: number, title: string) => `# Day ${day}: ${title}

## Part 1

`;

const createDir = (day: number, title: string) => {
  if (!day || !title) return;

  const dirName = day < 10 ? `day0${day}` : `day${day}`;
  const dirPath = path.join(process.cwd(), dirName);

  if (!fs.existsSync(dirPath)) {
    fs.mkdirSync(dirPath);
    console.log("Dir created:", dirName);
  }

  const indexPath = path.join(dirPath, "index.ts");
  if (!existsSync(indexPath)) {
    fs.writeFileSync(indexPath, getIndexContent(day, title));
    console.log("File created:", path.join(dirName, "index.ts"));
  }

  const readmePath = path.join(dirPath, "README.md");
  if (!existsSync(readmePath)) {
    fs.writeFileSync(readmePath, getReadmeContent(day, title));
    console.log("File created:", path.join(dirName, "README.md"));
  }

  const testInputPath = path.join(dirPath, "input-test.txt");
  if (!existsSync(testInputPath)) {
    fs.writeFileSync(testInputPath, "");
    console.log("File created:", path.join(dirName, "input-test.txt"));
  }

  // Add package script and sort package.json (pkg set doesn't respect order)
  execSync(`npm pkg set scripts."day:${day}"="npm run --silent execute -- ./${dirName}"`);
  execSync("npx --yes sort-package-json");
  console.log("Script added:", `npm run day:${day}`);
};

const init = async () => {
  const day = await digitPrompt("Day");
  const title = await textPrompt("Puzzle title");

  createDir(day, title);
};

(async () => {
  await init();
})();