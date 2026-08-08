import Link from "next/link";

export default function Navbar(){

return(

<nav className="flex items-center justify-between border-b border-slate-200 bg-white px-10 py-6 shadow-sm">

<Link
href="/"
className="text-3xl font-bold text-slate-900"
>
Redira
</Link>

<div className="flex gap-10">

<Link
href="/dashboard"
className="text-slate-700 hover:text-blue-600"
>
Dashboard
</Link>

<Link
href="/links"
className="text-slate-700 hover:text-blue-600"
>
Links
</Link>

<Link
href="/analytics"
className="text-slate-700 hover:text-blue-600"
>
Analytics
</Link>

</div>

</nav>

);

}