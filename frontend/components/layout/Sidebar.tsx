"use client";

import Link from "next/link";
import {LayoutDashboard,Link2,BarChart3} from "lucide-react";

export default function Sidebar(){

return(
<aside className="fixed left-0 top-0 h-screen w-72 bg-black border-r border-white/10 p-6 flex flex-col">

<h1 className="text-3xl font-bold text-white">
Redira
</h1>

<p className="text-sm text-slate-500 mt-1">
URL Analytics
</p>


<nav className="mt-10 space-y-2">

<LinkItem href="/dashboard" icon={<LayoutDashboard size={18}/>} text="Dashboard"/>

<LinkItem href="/links" icon={<Link2 size={18}/>} text="Links"/>

<LinkItem href="/analytics" icon={<BarChart3 size={18}/>} text="Analytics"/>

</nav>


<div className="mt-auto rounded-xl border border-white/10 bg-slate-950 p-4">

<p className="text-white font-medium">
Yakubu
</p>

<p className="text-sm text-slate-500">
Pro Account
</p>

</div>


</aside>
)

}


function LinkItem({
href,
icon,
text
}:{
href:string;
icon:any;
text:string;
}){

return(

<Link
href={href}
className="
flex
items-center
gap-3
rounded-xl
px-4
py-3
text-slate-400
hover:bg-white/5
hover:text-white
transition
"
>

{icon}

{text}

</Link>

)

}
