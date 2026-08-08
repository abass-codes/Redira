"use client";

import {useState} from "react";


export default function AnalyticsPreview(){

const [tab,setTab]=useState("Clicks");


const tabs=[
"Clicks",
"Devices",
"Locations"
];


return(

<section className="px-6 py-24">


<div className="mx-auto max-w-5xl rounded-3xl border border-slate-800 bg-slate-950 p-10">


<h2 className="text-3xl font-bold text-white">
Analytics Dashboard
</h2>


<p className="mt-3 text-slate-400">
Understand how your links perform.
</p>


<div className="mt-8 flex gap-3">


{tabs.map((item)=>(

<button
key={item}
onClick={()=>setTab(item)}
className={`rounded-lg px-5 py-2 ${
tab===item
?
"bg-blue-600 text-white"
:
"bg-slate-800 text-slate-300"
}`}
>

{item}

</button>

))}


</div>


<div className="mt-8 rounded-2xl bg-slate-900 p-8">


<p className="text-slate-400">
Showing {tab} analytics
</p>


<div className="mt-6 flex h-32 items-end gap-4">


<div className="h-12 w-full rounded bg-blue-500"/>

<div className="h-20 w-full rounded bg-purple-500"/>

<div className="h-28 w-full rounded bg-pink-500"/>


</div>


</div>


</div>


</section>

);

}