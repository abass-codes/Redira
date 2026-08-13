"use client";

import api from "@/lib/api";

export default function LinkActions({
id,
active,
refresh
}:{
id:string;
active:boolean;
refresh:()=>void;
}){


async function toggle(){

if(active){

await api.post(`/links/${id}/disable`);

}else{

await api.post(`/links/${id}/enable`);

}

refresh();

}



async function remove(){

await api.delete(`/links/${id}`);

refresh();

}



return(

<div className="flex gap-3 mt-5">


<button
onClick={toggle}
className="rounded-xl border border-slate-700 px-4 py-2 text-white hover:bg-slate-800"
>

{active ? "Disable" : "Enable"}

</button>


<button
onClick={remove}
className="rounded-xl bg-red-600 px-4 py-2 text-white hover:bg-red-500"
>

Delete

</button>


</div>

);

}
