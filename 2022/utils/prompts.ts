import prompts from "prompts";

export const digitPrompt = async (message: string) => {
  const input = await prompts({
    type: "number",
    name: "response",
    message,
  });

  return input.response as number;
};

export const textPrompt = async (message: string) => {
  const input = await prompts({
    type: "text",
    name: "response",
    message,
  });

  return input.response as string;
};
