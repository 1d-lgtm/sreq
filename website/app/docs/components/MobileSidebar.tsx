"use client";

import { useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Menu, X, ChevronDown } from "lucide-react";
import type { SidebarSection } from "../../lib/docs";

interface MobileSidebarProps {
  sections: SidebarSection[];
  basePath: string;
}

export default function MobileSidebar({
  sections,
  basePath,
}: MobileSidebarProps) {
  const [open, setOpen] = useState(false);
  const pathname = usePathname();

  return (
    <div className="lg:hidden">
      <button
        onClick={() => setOpen(true)}
        className="fixed top-4 left-4 z-30 p-2 rounded-md bg-neutral-900 border border-neutral-800 text-neutral-400 hover:text-white transition-colors"
        aria-label="Open menu"
      >
        <Menu className="w-5 h-5" />
      </button>

      {open && (
        <>
          <div
            className="fixed inset-0 bg-black/60 z-40"
            onClick={() => setOpen(false)}
          />
          <div className="fixed inset-y-0 left-0 w-[280px] bg-[#0a0a0a] border-r border-neutral-800 z-50 overflow-y-auto">
            <div className="flex items-center justify-between p-5">
              <div>
                <Link
                  href={basePath || "/"}
                  className="text-lg font-bold text-white"
                  onClick={() => setOpen(false)}
                >
                  sreq
                </Link>
                <span className="text-neutral-600 mx-2">|</span>
                <span className="text-sm text-neutral-400">docs</span>
              </div>
              <button
                onClick={() => setOpen(false)}
                className="p-1.5 rounded-md text-neutral-400 hover:text-white hover:bg-neutral-800 transition-colors"
                aria-label="Close menu"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <nav className="px-3 pb-8">
              {sections.map((section) => (
                <MobileGroup
                  key={section.heading || section.links[0]?.href}
                  section={section}
                  pathname={pathname}
                  basePath={basePath}
                  onNavigate={() => setOpen(false)}
                />
              ))}
            </nav>
          </div>
        </>
      )}
    </div>
  );
}

function MobileGroup({
  section,
  pathname,
  basePath,
  onNavigate,
}: {
  section: SidebarSection;
  pathname: string;
  basePath: string;
  onNavigate: () => void;
}) {
  const [expanded, setExpanded] = useState(true);

  return (
    <div className="mb-1">
      {section.heading ? (
        <button
          onClick={() => setExpanded(!expanded)}
          className="flex items-center justify-between w-full px-3 py-2 mt-4 text-xs font-semibold uppercase tracking-wider text-neutral-400"
        >
          {section.heading}
          <ChevronDown
            className={`w-3.5 h-3.5 transition-transform ${expanded ? "" : "-rotate-90"}`}
          />
        </button>
      ) : null}

      {expanded && (
        <ul className="space-y-0.5">
          {section.links.map((link) => {
            const fullHref = `${basePath}${link.href}`;
            const isActive =
              pathname === fullHref || pathname === `${fullHref}/`;

            return (
              <li key={link.href}>
                <Link
                  href={fullHref}
                  onClick={onNavigate}
                  className={`block px-3 py-2 text-sm rounded-md transition-colors ${
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
