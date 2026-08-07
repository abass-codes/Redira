"use client";

import Link from "next/link";
import { logout } from "@/lib/api";

export default function Sidebar() {
  return (
    <aside className="w-64 min-h-screen bg-zinc-950 border-r border-zinc-800 p-6 text-white">
      <h1 className="text-3xl font-bold mb-10">
        Redira
      </h1>

      <nav className="space-y-5">
        <Link
          href="/dashboard"
          className="block hover:text-blue-400"
        >
          Dashboard
        </Link>

        <Link
          href="/links"
          className="block hover:text-blue-400"
        >
          Links
        </Link>

        <button
          onClick={logout}
          className="text-red-400 hover:text-red-300"
        >
          Logout
        </button>
      </nav>
    </aside>
  );
}