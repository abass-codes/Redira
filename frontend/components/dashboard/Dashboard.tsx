"use client";

import useDashboard from "@/hooks/useDashboard";
import StatsCard from "./StatsCard";

export default function Dashboard() {
  const { dashboard, loading } = useDashboard();

  if (loading) {
    return (
      <div className="text-slate-400">
        Loading dashboard...
      </div>
    );
  }

  return (
    <div className="space-y-10">

      <div>
        <h1 className="text-4xl font-bold text-white">
          Dashboard
        </h1>

        <p className="mt-2 text-slate-400">
          Monitor your links and performance.
        </p>
      </div>


      <div className="grid gap-6 md:grid-cols-3">

        <StatsCard
          title="Total Links"
          value={dashboard?.TotalLinks ?? 0}
          description="Created links"
        />

        <StatsCard
          title="Total Clicks"
          value={dashboard?.TotalClicks ?? 0}
          description="Total visits"
        />

        <StatsCard
          title="Active Links"
          value={dashboard?.ActiveLinks ?? 0}
          description="Currently active"
        />

      </div>


      <div className="rounded-2xl border border-slate-800 bg-slate-950 p-8">

        <h2 className="text-2xl font-bold text-white">
          Performance Overview
        </h2>

        <p className="mt-3 text-slate-400">
          Your link activity will appear here.
        </p>

        <div className="mt-8 flex h-48 items-center justify-center rounded-xl bg-slate-900">

          <p className="text-slate-500">
            Analytics visualization
          </p>

        </div>

      </div>

    </div>
  );
}