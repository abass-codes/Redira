"use client";

import {useEffect,useState} from "react";
import api from "@/lib/api";

import StatsCard from "@/components/dashboard/StatsCard";
import ActivityChart from "@/components/dashboard/ActivityChart";
import SystemStatus from "@/components/dashboard/SystemStatus";


export default function Dashboard(){


const [stats,setStats]=useState({
TotalLinks:0,
TotalClicks:0,
ActiveLinks:0
});


useEffect(()=>{

api.get("/dashboard")
.then(res=>setStats(res.data));

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


<div>

<h1 className="
text-5xl
font-bold
">
Dashboard
</h1>


<p className="
mt-3
text-slate-400
">
Monitor your links and performance.
</p>

</div>



<div className="
mt-10
grid
grid-cols-3
gap-6
">


<StatsCard
title="Total Links"
value={stats.TotalLinks}
subtitle="Created links"
/>


<StatsCard
title="Total Clicks"
value={stats.TotalClicks}
subtitle="Total visits"
/>


<StatsCard
title="Active Links"
value={stats.ActiveLinks}
subtitle="Currently active"
/>


</div>



<div className="mt-8">

<ActivityChart/>

</div>



<div className="
mt-8
grid
grid-cols-2
gap-6
">


<div className="
rounded-2xl
border
border-white/10
bg-slate-950
p-6
">


<h2 className="text-xl font-bold">
Quick Insights
</h2>


<div className="mt-5 space-y-3 text-slate-400">

<p>
🔗 {stats.TotalLinks} links created
</p>

<p>
📈 {stats.TotalClicks} total visits
</p>

</div>


</div>


<SystemStatus/>


</div>



</div>


</main>

)

}
