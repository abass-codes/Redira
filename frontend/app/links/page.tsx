"use client";

import CreateLink from "@/components/links/CreateLink";
import useAuth from "@/hooks/useAuth";

export default function Links() {
  useAuth();

  return (
    <div className="p-10 text-white">
      <h1 className="text-4xl font-bold mb-8">
        Links
      </h1>

      <CreateLink />
    </div>
  );
}