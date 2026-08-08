"use client";

import {useState} from "react";


export default function UrlShortenerDemo(){

const [url,setUrl]=useState("");

const [short,setShort]=useState("");

const [copied,setCopied]=useState(false);


function shorten(){

if(!url)return;

setShort("redira.dev/a8K29x");

}


function copy(){

navigator.clipboard.writeText(short);

setCopied(true);

setTimeout(()=>{

setCopied(false);

},2000);

}


return(

<div className="mx-auto max-w-3xl">


<div className="flex flex-col gap-3 sm:flex-row">


<input
value={url}
onChange={(e)=>setUrl(e.target.value)}
placeholder="Paste your long URL..."
className="flex-1 rounded-xl border border-slate-700 bg-slate-900 px-5 py-4 text-white outline-none focus:border-blue-500"
/>


<button
onClick={shorten}
className="rounded-xl bg-blue-600 px-8 py-4 font-semibold text-white hover:bg-blue-700"
>
Shorten
</button>


</div>


{short && (

<div className="mt-5 flex items-center justify-between rounded-xl border border-slate-700 bg-slate-900 p-5">


<p className="text-blue-400">
{short}
</p>


<button
onClick={copy}
className="text-slate-300 hover:text-white"
>
{copied?"✓ Copied":"Copy"}
</button>


</div>

)}


</div>

);

}