import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Redira | Shorten Links. Track Performance.",
  description:
    "Redira is a modern URL shortening platform with analytics and performance tracking.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}