import Link from "next/link";

export default function Home(){

return(
<main className="min-h-screen flex flex-col items-center justify-center bg-slate-50">

<h1 className="text-7xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
Redira
</h1>

<p className="mt-6 text-xl text-slate-600">
A production-ready URL shortening platform with analytics and performance tracking.
</p>

<div className="mt-10 flex gap-5">

<Link
href="/login"
className="rounded-xl bg-blue-600 px-8 py-4 font-semibold text-white shadow-lg shadow-blue-200 hover:bg-blue-700"
>
Login
</Link>

<Link
href="/register"
className="rounded-xl border border-slate-300 bg-white px-8 py-4 font-semibold text-slate-900 hover:bg-slate-100"
>
Create Account
</Link>

</div>

<div className="mt-24 grid grid-cols-1 gap-6 md:grid-cols-3">

<div className="w-72 rounded-2xl border border-slate-200 bg-white p-8 text-center shadow-sm">

<h2 className="text-2xl font-bold text-slate-900">
Fast
</h2>

<p className="mt-3 text-slate-600">
Instant redirects
</p>

</div>


<div className="w-72 rounded-2xl border border-slate-200 bg-white p-8 text-center shadow-sm">

<h2 className="text-2xl font-bold text-slate-900">
Analytics
</h2>

<p className="mt-3 text-slate-600">
Track every click
</p>

</div>


<div className="w-72 rounded-2xl border border-slate-200 bg-white p-8 text-center shadow-sm">

<h2 className="text-2xl font-bold text-slate-900">
Reliable
</h2>

<p className="mt-3 text-slate-600">
Production infrastructure
</p>

</div>

</div>

</main>
);

}