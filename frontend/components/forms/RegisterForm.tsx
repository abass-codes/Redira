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
onClick={register}
disabled={loading}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 font-semibold text-white hover:bg-blue-700"
>

{loading?"Creating...":"Register"}

</button>

</>

);

}