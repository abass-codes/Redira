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

setError("Invalid email or password.");

}finally{

setLoading(false);

}

}

return(

<div className="w-full max-w-md rounded-3xl border border-slate-200 bg-white p-10 shadow-xl">

<div className="text-center">

<h1 className="text-3xl font-bold tracking-tight text-slate-900">
REDIRA
</h1>

<p className="mt-6 text-2xl font-semibold text-slate-900">
Welcome back
</p>

<p className="mt-2 text-sm text-slate-500">
Sign in to manage your links, redirects, and analytics.
</p>

</div>

{error&&(
<p className="mt-6 text-center text-sm text-red-500">
{error}
</p>
)}

<input
className="mt-8 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20"
placeholder="Email"
value={email}
onChange={(e)=>setEmail(e.target.value)}
/>

<input
className="mt-4 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20"
placeholder="Password"
type="password"
value={password}
onChange={(e)=>setPassword(e.target.value)}
/>

<button
onClick={login}
disabled={loading}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 font-semibold text-white transition hover:bg-blue-700 disabled:opacity-50"
>

{loading?"Signing in...":"Sign in"}

</button>

<p className="mt-6 text-center text-sm text-slate-500">
Don't have an account?{" "}
<a href="/register" className="font-medium text-blue-600 hover:text-blue-700">
Create account
</a>
</p>

</div>

);

}
