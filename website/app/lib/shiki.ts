import { createHighlighter, type Highlighter } from "shiki";

let highlighter: Highlighter | null = null;

export async function getHighlighter(): Promise<Highlighter> {
  if (!highlighter) {
    highlighter = await createHighlighter({
      themes: ["github-dark"],
      langs: ["bash", "yaml", "json", "go", "toml", "shell", "plaintext"],
    });
  }
  return highlighter;
}
