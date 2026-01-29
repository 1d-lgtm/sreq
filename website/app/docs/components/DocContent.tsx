"use client";

import { useEffect, useRef } from "react";
import { createRoot } from "react-dom/client";
import CopyCodeButton from "./CopyCodeButton";

interface DocContentProps {
  html: string;
}

export default function DocContent({ html }: DocContentProps) {
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!containerRef.current) return;

    const preBlocks = containerRef.current.querySelectorAll("pre");
    const roots: ReturnType<typeof createRoot>[] = [];

    preBlocks.forEach((pre) => {
      // Add group class for hover effect
      pre.classList.add("group");
      pre.style.position = "relative";

      // Extract text content for copy
      const code = pre.textContent || "";

      // Create mount point for copy button
      const mount = document.createElement("div");
      mount.className = "absolute top-0 right-0";
      pre.appendChild(mount);

      const root = createRoot(mount);
      root.render(<CopyCodeButton code={code} />);
      roots.push(root);
    });

    return () => {
      roots.forEach((root) => root.unmount());
    };
  }, [html]);

  return (
    <div
      ref={containerRef}
      className="prose prose-invert prose-emerald max-w-none prose-headings:tracking-tight prose-h1:text-3xl prose-h1:font-bold prose-h2:text-xl prose-h2:border-b prose-h2:border-neutral-800 prose-h2:pb-2 prose-p:text-neutral-400 prose-a:text-emerald-400 prose-a:no-underline hover:prose-a:text-emerald-300 prose-strong:text-neutral-200 prose-code:text-emerald-400 prose-code:bg-neutral-800 prose-code:px-1.5 prose-code:py-0.5 prose-code:rounded prose-code:font-normal prose-code:before:content-none prose-code:after:content-none prose-pre:bg-[#171717] prose-pre:border prose-pre:border-neutral-800 prose-pre:rounded-lg prose-th:text-neutral-200 prose-td:text-neutral-400 prose-li:text-neutral-400 prose-blockquote:border-emerald-500 prose-blockquote:bg-neutral-900 prose-blockquote:rounded-r-lg prose-blockquote:text-neutral-400 prose-hr:border-neutral-800"
      dangerouslySetInnerHTML={{ __html: html }}
    />
  );
}
