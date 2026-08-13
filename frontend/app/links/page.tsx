"use client";

import CreateLink from "@/components/links/CreateLink";
import LinkTable from "@/components/links/LinkTable";

export default function LinksPage(){

return(

<main className="min-h-screen bg-black text-white p-8 flex justify-center">

<div className="w-full max-w-7xl">


<div className="mb-8">

<h1 className="text-5xl font-bold">
Links
</h1>

<p className="mt-3 text-slate-400">
Create, manage, and track your short links.
</p>

</div>


<CreateLink/>


<div className="mt-8">

<LinkTable/>

</div>


</div>

</main>

)

}
