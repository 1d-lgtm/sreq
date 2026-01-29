import type { Metadata } from "next";
import { parseSidebar } from "../lib/docs";
import Sidebar from "./components/Sidebar";
import MobileSidebar from "./components/MobileSidebar";

export const metadata: Metadata = {
  title: "sreq docs",
  description:
    "Documentation for sreq — service-aware API client with automatic credential resolution.",
};

const basePath = process.env.NODE_ENV === "production" ? "/sreq" : "";

export default function DocsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const sections = parseSidebar();

  return (
    <div className="min-h-screen bg-[#0a0a0a]">
      <Sidebar sections={sections} basePath={basePath} />
      <MobileSidebar sections={sections} basePath={basePath} />
      <main className="lg:ml-[260px] xl:mr-[200px] min-h-screen">
        <div className="max-w-[750px] mx-auto px-6 py-16 lg:py-12">
          {children}
        </div>
      </main>
    </div>
  );
}
