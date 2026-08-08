"use client";

import useLinks from "@/hooks/useLinks";

export default function LinkTable(){

const {links,loading}=useLinks();

if(loading){

return <p className="text-slate-400">Loading links...</p>;

}

return(

<div className="rounded-2xl border border-slate-800 bg-slate-950 p-8">

<h2 className="text-2xl font-bold text-white">
Your Links
</h2>

<div className="mt-6 space-y-4">

{links.length===0&&(
<p className="text-slate-400">
No links created yet.
</p>
)}

{links.map((link)=>(

<div
key={link.id}
className="rounded-xl border border-slate-800 bg-slate-900 p-5"
>

<p className="truncate text-white">
{link.original_url}
</p>

<div className="mt-3 flex justify-between">

<p className="text-blue-400">
{link.short_code}
</p>

<p className="text-slate-400">
{link.clicks??0} clicks
</p>

</div>

</div>

))}

</div>

</div>

);

}