"use client";

import { useEffect, useState } from "react";
import type { TocItem } from "../../lib/docs";

interface TableOfContentsProps {
  items: TocItem[];
}

export default function TableOfContents({ items }: TableOfContentsProps) {
  const [activeId, setActiveId] = useState<string>("");

  useEffect(() => {
    if (items.length === 0) return;

    const observer = new IntersectionObserver(
      (entries) => {
        for (const entry of entries) {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id);
          }
        }
      },
      { rootMargin: "-80px 0px -70% 0px", threshold: 0 }
    );

    for (const item of items) {
      const el = document.getElementById(item.id);
      if (el) observer.observe(el);
    }

    return () => observer.disconnect();
  }, [items]);

  if (items.length === 0) return null;

  return (
    <aside className="hidden xl:block fixed right-0 top-0 w-[200px] h-screen pt-20 pr-6 overflow-y-auto">
      <p className="text-xs font-semibold uppercase tracking-wider text-neutral-500 mb-3">
        On this page
      </p>
      <ul className="space-y-1.5">
        {items.map((item) => (
          <li key={item.id}>
            <a
              href={`#${item.id}`}
              className={`block text-[13px] leading-snug transition-colors ${
                item.level === 3 ? "pl-3" : ""
              } ${
                activeId === item.id
                  ? "text-emerald-400"
                  : "text-neutral-500 hover:text-neutral-300"
              }`}
            >
              {item.text}
            </a>
          </li>
        ))}
      </ul>
    </aside>
  );
}
