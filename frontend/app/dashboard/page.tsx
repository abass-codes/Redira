"use client";

import { useEffect, useState } from "react";
import api from "@/lib/api";
import Card from "@/components/dashboard/Card";
import useAuth from "@/hooks/useAuth";

export default function Dashboard() {
  useAuth();

  const [data, setData] = useState<any>(null);

  useEffect(() => {
    api
      .get("/dashboard")
      .then((res) => setData(res.data));
  }, []);

  if (!data) {
    return (
      <div className="p-10 text-white">
        Loading...
      </div>
    );
  }

  return (
    <div className="p-10 text-white">
      <h1 className="text-4xl font-bold mb-8">
        Dashboard
      </h1>

      <div className="grid grid-cols-3 gap-6">
        <Card
          title="Total Links"
          value={data.TotalLinks}
        />

        <Card
          title="Total Clicks"
          value={data.TotalClicks}
        />

        <Card
          title="Active Links"
          value={data.ActiveLinks}
        />
      </div>
    </div>
  );
}