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

const response=await api.post("/api/v1/auth/login",{
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

<>
{error&&(
<p className="mt-4 text-red-500">
{error}
</p>
)}

<input
className="mt-6 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none focus:border-blue-500"
placeholder="Email"
value={email}
onChange={(e)=>setEmail(e.target.value)}
/>

<input
className="mt-4 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none focus:border-blue-500"
placeholder="Password"
type="password"
value={password}
onChange={(e)=>setPassword(e.target.value)}
/>

<button
onClick={login}
disabled={loading}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 font-semibold text-white hover:bg-blue-700"
>

{loading?"Logging in...":"Login"}

</button>

</>
);

}