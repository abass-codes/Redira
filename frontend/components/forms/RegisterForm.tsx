"use client";

import {useState} from "react";
import api from "@/lib/api";
import {useRouter} from "next/navigation";

export default function RegisterForm(){

const router=useRouter();

const [email,setEmail]=useState("");
const [password,setPassword]=useState("");
const [loading,setLoading]=useState(false);
const [error,setError]=useState("");

async function register(){

try{

setLoading(true);
setError("");

await api.post("/auth/register",{
email,
password
});

router.push("/login");

}catch{

setError("Unable to create account. Please try again.");

}finally{

setLoading(false);

}

}

return(

<div className="w-full max-w-md rounded-3xl border border-slate-200 bg-white/90 p-10 shadow-2xl backdrop-blur">

<div className="text-center">

<p className="text-xl font-semibold tracking-[0.3em] text-blue-600">
REDIRA
</p>

<h1 className="mt-6 text-3xl font-bold tracking-tight text-slate-900">
Create your account
</h1>

<p className="mt-3 text-sm text-slate-500">
Start managing short links, redirects, and analytics.

</p>

</div>

{error&&(
<p className="mt-6 rounded-xl bg-red-50 px-4 py-3 text-center text-sm text-red-600">
{error}
</p>
)}

<div className="mt-8">

<label className="text-sm font-medium text-slate-700">
Email address
</label>

<input
className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none transition focus:border-blue-500 focus:ring-4 focus:ring-blue-500/10"
placeholder="you@example.com"
value={email}
onChange={(e)=>setEmail(e.target.value)}
/>

</div>

<div className="mt-5">

<label className="text-sm font-medium text-slate-700">
Password
</label>

<input
className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none transition focus:border-blue-500 focus:ring-4 focus:ring-blue-500/10"
placeholder="Create a password"
type="password"
value={password}
onChange={(e)=>setPassword(e.target.value)}
/>

</div>

<button
onClick={register}
disabled={loading}
className="mt-7 w-full rounded-xl bg-blue-600 py-3.5 font-semibold text-white transition hover:bg-blue-700 disabled:opacity-50"
>

{loading?"Creating account...":"Create account"}

</button>

<p className="mt-6 text-center text-sm text-slate-500">
Already have an account?{" "}
<a
href="/login"
className="font-medium text-blue-600 hover:text-blue-700"
>
Sign in
</a>
</p>

<p className="mt-8 text-center text-xs text-slate-400">
Protected by secure authentication
</p>

</div>

);

}