"use client";

import {useState} from "react";
import api from "@/lib/api";
import {useRouter} from "next/navigation";

export default function RegisterForm(){

const router=useRouter();

const [email,setEmail]=useState("");
const [password,setPassword]=useState("");

async function register(){

await api.post(
"/api/v1/auth/register",
{
email,
password
}
);

router.push("/login");

}

return(

<>

<input
className="mt-6 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none focus:border-blue-500"
placeholder="Email"
onChange={(e)=>setEmail(e.target.value)}
/>

<input
className="mt-4 w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 outline-none focus:border-blue-500"
placeholder="Password"
type="password"
onChange={(e)=>setPassword(e.target.value)}
/>

<button
onClick={register}
className="mt-6 w-full rounded-xl bg-blue-600 py-3 font-semibold text-white hover:bg-blue-700"
>

Register

</button>

</>

);

}