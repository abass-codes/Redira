"use client";

import {useEffect,useState} from "react";
import api from "@/lib/api";


export default function AnalyticsPage(){

const [links,setLinks]=useState<any[]>([]);


useEffect(()=>{

async function load(){

const response=await api.get("/links");

setLinks(response.data.links ?? []);

}

load();

},[]);



return(

<main className="
min-h-screen
bg-black
text-white
p-8
flex
justify-center
">

<div className="w-full max-w-7xl">


<div className="mb-10">

<h1 className="text-5xl font-bold">
Analytics
</h1>

<p className="mt-3 text-slate-400">
Track link activity and performance history.
</p>

</div>



<div className="
rounded-2xl
border
border-white/10
bg-slate-950
overflow-hidden
">


<div className="
grid
grid-cols-12
px-6
py-4
text-xs
uppercase
text-slate-500
border-b
border-white/10
">


<div className="col-span-5">
Link
</div>


<div className="col-span-1">
Clicks
</div>


<div className="col-span-3">
Created
</div>


<div className="col-span-3">
Last Clicked
</div>


</div>



{
links.map(link=>(

<div
key={link.ID}
className="
grid
grid-cols-12
items-center
px-6
py-5
border-b
border-white/5
hover:bg-white/[0.03]
"
>


<div className="col-span-5">

<p className="text-white truncate">
{link.OriginalUrl}
</p>

<p className="text-blue-400 text-sm">
redira/{link.ShortCode}
</p>

</div>



<div className="col-span-1 text-white font-bold text-xl">

{link.ClickCount ?? 0}

</div>



<div className="col-span-3 text-slate-400">

{new Date(link.CreatedAt).toLocaleString()}

</div>



<div className="col-span-3 text-slate-400">

{
link.LastClickedAt
?
new Date(link.LastClickedAt).toLocaleString()
:
"Never"
}

</div>



</div>

))

}



</div>


</div>

</main>

)

}
