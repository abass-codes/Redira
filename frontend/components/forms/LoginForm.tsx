"use client";

import {useState} from "react";
import api from "@/lib/api";
import {useRouter} from "next/navigation";
import {saveToken} from "@/lib/auth";

export default function LoginForm(){

const router=useRouter();

const [email,setEmail]=useState("");
const [password,setPassword]=useState("");
const [loading,setLoading]=useState(false);
const [error,setError]=useState("");

async function login(){

try{

setLoading(true);
setError("");

const response=await api.post("/auth/login",{
email,
password
});

saveToken(response.data.token);

router.push("/dashboard");

}catch{

setError("Unable to sign in. Check your email and password and try again.");

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
Welcome back
</h1>

<p className="mt-3 text-sm text-slate-500">
Sign in to manage links, redirects, and analytics.

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
placeholder="••••••••"
type="password"
value={password}
onChange={(e)=>setPassword(e.target.value)}
/>

</div>

<button
onClick={login}
disabled={loading}
className="mt-7 w-full rounded-xl bg-blue-600 py-3.5 font-semibold text-white transition hover:bg-blue-700 disabled:opacity-50"
>

{loading?"Signing in...":"Sign in"}

</button>

<p className="mt-6 text-center text-sm text-slate-500">
New to Redira?{" "}
<a
href="/register"
className="font-medium text-blue-600 hover:text-blue-700"
>
Create an account
</a>
</p>

<p className="mt-8 text-center text-xs text-slate-400">
Protected by secure authentication
</p>

</div>

);

}