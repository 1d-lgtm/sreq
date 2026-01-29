"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronDown } from "lucide-react";
import { useState } from "react";
import type { SidebarSection } from "../../lib/docs";

interface SidebarProps {
  sections: SidebarSection[];
  basePath: string;
}

export default function Sidebar({ sections, basePath }: SidebarProps) {
  const pathname = usePathname();

  return (
    <aside className="hidden lg:block fixed left-0 top-0 bottom-0 w-[260px] border-r border-neutral-800 bg-[#0a0a0a] overflow-y-auto z-20">
      <div className="p-5 pb-3">
        <Link
          href={basePath || "/"}
          className="text-lg font-bold tracking-tight text-white hover:text-emerald-400 transition-colors"
        >
          sreq
        </Link>
        <span className="text-neutral-600 mx-2">|</span>
        <Link
          href={`${basePath}/docs`}
          className="text-sm text-neutral-400 hover:text-neutral-200 transition-colors"
        >
          docs
        </Link>
      </div>

      <nav className="px-3 pb-8">
        {sections.map((section) => (
          <SidebarGroup
            key={section.heading || section.links[0]?.href}
            section={section}
            pathname={pathname}
            basePath={basePath}
          />
        ))}
      </nav>
    </aside>
  );
}

function SidebarGroup({
  section,
  pathname,
  basePath,
}: {
  section: SidebarSection;
  pathname: string;
  basePath: string;
}) {
  const hasActiveChild = section.links.some(
    (link) => pathname === `${basePath}${link.href}` || pathname === `${basePath}${link.href}/`
  );
  const [open, setOpen] = useState(hasActiveChild || !section.heading);

  return (
    <div className="mb-1">
      {section.heading ? (
        <button
          onClick={() => setOpen(!open)}
          className="flex items-center justify-between w-full px-3 py-2 mt-4 text-xs font-semibold uppercase tracking-wider text-neutral-400 hover:text-neutral-200 transition-colors"
        >
          {section.heading}
          <ChevronDown
            className={`w-3.5 h-3.5 transition-transform ${open ? "" : "-rotate-90"}`}
          />
        </button>
      ) : null}

      {open && (
        <ul className="space-y-0.5">
          {section.links.map((link) => {
            const fullHref = `${basePath}${link.href}`;
            const isActive = pathname === fullHref || pathname === `${fullHref}/`;

            return (
              <li key={link.href}>
                <Link
                  href={fullHref}
                  className={`block px-3 py-1.5 text-sm rounded-md transition-colors ${
                    isActive
                      ? "text-emerald-400 bg-emerald-500/10 font-medium"
                      : "text-neutral-400 hover:text-neutral-200 hover:bg-neutral-800/50"
                  }`}
                >
                  {link.label}
                </Link>
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}
