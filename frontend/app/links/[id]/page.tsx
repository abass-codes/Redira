"use client";

import {useEffect,useState} from "react";
import {useParams} from "next/navigation";
import api from "@/lib/api";


export default function LinkAnalytics(){

const params = useParams();

const id = params.id as string;


const [data,setData]=useState<any>(null);


useEffect(()=>{

async function load(){

const response = await api.get(
`/analytics/links/${id}`
);

setData(response.data);

}

if(id){
load();
}

},[id]);



if(!data){

return(

<main className="min-h-screen bg-black text-white p-10">

<p className="text-slate-400">
Loading analytics...
</p>

</main>

)

}



return(

<main className="min-h-screen bg-black text-white p-10">


<h1 className="text-5xl font-bold">
Link Analytics
</h1>



<div className="mt-10 rounded-3xl border border-slate-800 bg-slate-950 p-8">


<h2 className="text-2xl font-bold">
Click Events
</h2>


<p className="mt-5 text-5xl font-bold">
{data.length}
</p>


<p className="text-slate-400">
total clicks
</p>


</div>



<div className="mt-10 rounded-3xl border border-slate-800 bg-slate-950 p-8">


<h2 className="text-2xl font-bold">
Recent Activity
</h2>



<div className="mt-6 space-y-4">


{
data.map((event:any,index:number)=>(


<div
key={index}
className="rounded-xl bg-slate-900 p-5"
>


<p className="text-white">
Click #{index+1}
</p>


<p className="mt-2 text-slate-400">
{event.created_at}
</p>


</div>


))

}


</div>


</div>


</main>

)

}
