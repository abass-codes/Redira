"use client";

import {useState} from "react";
import api from "@/lib/api";

export default function CreateLinkForm(){

const [url,setUrl]=useState("");
const [shortUrl,setShortUrl]=useState("");
const [loading,setLoading]=useState(false);

async function createLink(){

if(!url)return;

try{

setLoading(true);

const response=await api.post("/links",{url});

setShortUrl(response.data.short_url);

setUrl("");

}finally{

setLoading(false);

}

}

return(

<div className="rounded-2xl border border-slate-800 bg-slate-950 p-8">

<h2 className="text-2xl font-bold text-white">
Create Short Link
</h2>

<p className="mt-2 text-slate-400">
Convert long URLs into shareable links.
</p>

<div className="mt-6 flex flex-col gap-4">

<input
value={url}
onChange={(e)=>setUrl(e.target.value)}
placeholder="https://example.com"
className="w-full rounded-xl border border-slate-700 bg-slate-900 px-5 py-3 text-white outline-none focus:border-blue-500"
/>

<button
onClick={createLink}
disabled={loading}
className="w-full rounded-xl bg-blue-600 px-6 py-3 font-semibold text-white hover:bg-blue-700"
>
{loading?"Creating...":"Create Link"}
</button>

</div>

{shortUrl&&(
<div className="mt-6 rounded-xl border border-blue-900 bg-blue-950 p-4">
<p className="text-sm text-slate-300">
Your shortened URL
</p>

<p className="mt-2 text-blue-400">
{shortUrl}
</p>
</div>
)}

</div>

);

}