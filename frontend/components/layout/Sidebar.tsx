"use client";

import Link from "next/link";
import {useRouter} from "next/navigation";
import {removeToken} from "@/lib/auth";

export default function Sidebar(){

const router=useRouter();

function logout(){

removeToken();

router.push("/login");

}

return(

<aside className="min-h-screen w-64 border-r border-slate-800 bg-slate-950 p-6">

<h1 className="text-2xl font-bold text-white">
Redira
</h1>

<nav className="mt-8 space-y-4">

<Link
href="/dashboard"
className="block text-slate-300 hover:text-white"
>
Dashboard
</Link>

<Link
href="/links"
className="block text-slate-300 hover:text-white"
>
Links
</Link>

<Link
href="/analytics"
className="block text-slate-300 hover:text-white"
>
Analytics
</Link>

<button
onClick={logout}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 text-white hover:bg-blue-700"
>
Logout
</button>

</nav>

</aside>

);

}