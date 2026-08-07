"use client";

import { useState } from "react";
import api from "@/lib/api";

export default function CreateLink() {
  const [url, setUrl] = useState("");

  async function createLink() {
    await api.post("/links", {
      url,
    });

    alert("Link created");
    setUrl("");
  }

  return (
    <div className="bg-zinc-900 rounded-xl p-6">
      <input
        className="w-full bg-black border border-zinc-700 rounded p-3"
        placeholder="Enter URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />

      <button
        onClick={createLink}
        className="mt-4 bg-blue-600 px-5 py-2 rounded"
      >
        Create Link
      </button>
    </div>
  );
}