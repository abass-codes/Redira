"use client";

import { useEffect,useState } from "react";
import api from "@/lib/api";
import Card from "@/components/dashboard/Card";
import {DashboardSummary} from "@/types/dashboard";


export default function Dashboard(){

const [data,setData]=useState<DashboardSummary|null>(null);


useEffect(()=>{

api.get("/dashboard")
.then(res=>{
setData(res.data);
});

},[]);


if(!data){

return (
<div className="min-h-screen bg-black text-white flex items-center justify-center">
Loading...
</div>
);

}


return (

<main className="min-h-screen bg-black p-10 text-white">


<h1 className="text-4xl font-bold mb-8">
Redira Dashboard
</h1>


<div className="grid grid-cols-3 gap-6">


<Card
title="Total Links"
value={data.TotalLinks}
/>


<Card
title="Total Clicks"
value={data.TotalClicks}
/>


<Card
title="Active Links"
value={data.ActiveLinks}
/>


</div>


</main>

);

}