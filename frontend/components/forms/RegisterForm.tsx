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

setError("Unable to create account.");

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
Create your account
</p>

<p className="mt-2 text-sm text-slate-500">
Start managing short links, redirects, and analytics.
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
onClick={register}
disabled={loading}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 font-semibold text-white transition hover:bg-blue-700 disabled:opacity-50"
>

{loading?"Creating account...":"Create account"}

</button>

<p className="mt-6 text-center text-sm text-slate-500">
Already have an account?{" "}
<a href="/login" className="font-medium text-blue-600 hover:text-blue-700">
Sign in
</a>
</p>

</div>

);

}
