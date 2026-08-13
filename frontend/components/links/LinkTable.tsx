"use client";

import useLinks from "@/hooks/useLinks";
import Link from "next/link";
import api from "@/lib/api";


export default function LinkTable(){

const {links,loading,refresh}=useLinks();

async function deleteLink(id:string){

const confirmDelete = window.confirm(
"Delete this link?"
);

if(!confirmDelete) return;

await api.delete(`/links/${id}`);

refresh();

}


if(loading)
return <p className="text-slate-400">Loading...</p>;


return(

<div className="
w-full
max-w-7xl
mx-auto
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

<div className="col-span-6">
URL
</div>


<div className="col-span-4">
Short Link
</div>


<div className="col-span-1">
Clicks
</div>

<div className="col-span-1 text-right">
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
py-4
border-b
border-white/5
hover:bg-white/[0.03]
transition
"
>


<div className="
col-span-6
truncate
text-white
">

{link.OriginalUrl}

</div>



<div className="
col-span-4
">

<Link
href={`http://localhost:8080/r/${link.ShortCode}`}
target="_blank"
className="
text-blue-400
hover:text-blue-300
transition
"
>

redira/{link.ShortCode}

</Link>


</div>



<div className="
col-span-1
text-white
font-bold
">

{link.ClickCount ?? 0}

</div>


<div className="col-span-1 text-right">

<button
onClick={()=>deleteLink(link.ID)}
className="
rounded-lg
border
border-red-500/30
px-3
py-1
text-sm
text-red-400
hover:bg-red-500/10
"
>
Delete
</button>

</div>



</div>


))

}


</div>

)

}
