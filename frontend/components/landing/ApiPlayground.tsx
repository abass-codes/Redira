"use client";

import {useState} from "react";


export default function ApiPlayground(){

const [copied,setCopied]=useState(false);


const code=
`curl -X POST https://api.redira.dev/v1/links`;


function copy(){

navigator.clipboard.writeText(code);

setCopied(true);

setTimeout(()=>{

setCopied(false);

},2000);

}


return(

<section className="px-6 pb-24">


<div className="mx-auto max-w-5xl rounded-3xl border border-slate-800 bg-slate-950 p-10">


<h2 className="text-3xl font-bold text-white">
Developer API
</h2>


<p className="mt-3 text-slate-400">
Create and manage links programmatically.
</p>


<pre className="mt-8 rounded-xl bg-black p-6 text-blue-400">
{code}
</pre>


<button
onClick={copy}
className="mt-6 rounded-xl bg-blue-600 px-6 py-3 font-semibold text-white hover:bg-blue-700"
>
{copied?"✓ Copied":"Copy Code"}
</button>


</div>


</section>

);

}