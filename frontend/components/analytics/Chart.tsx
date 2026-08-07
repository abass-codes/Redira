"use client";

import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

export default function Chart({
  data,
}: {
  data: any[];
}) {
  return (
    <ResponsiveContainer
      width="100%"
      height={300}
    >
      <LineChart data={data}>
        <XAxis dataKey="Date" />
        <YAxis />
        <Tooltip />

        <Line
          type="monotone"
          dataKey="Clicks"
        />
      </LineChart>
    </ResponsiveContainer>
  );
}