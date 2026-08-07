import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen bg-black text-white flex items-center justify-center">

      <div className="text-center">

        <h1 className="text-6xl font-bold">
          Redira
        </h1>

        <p className="mt-4 text-gray-400">
          Shorten links. Track clicks. Understand traffic.
        </p>


        <div className="mt-8 flex gap-4 justify-center">

          <Link
            href="/login"
            className="rounded-lg bg-blue-600 px-6 py-3"
          >
            Login
          </Link>


          <Link
            href="/dashboard"
            className="rounded-lg border border-gray-700 px-6 py-3"
          >
            Dashboard
          </Link>

        </div>

      </div>

    </main>
  );
}