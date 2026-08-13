export default function SystemStatus(){

return(

<div className="
rounded-2xl
border
border-white/10
bg-slate-950
p-6
">

<h2 className="text-xl font-bold text-white">
System Status
</h2>


<div className="mt-6 space-y-5">


<div className="flex justify-between">
<span className="text-slate-400">
API
</span>

<span className="text-green-400">
Operational
</span>

</div>


<div className="flex justify-between">
<span className="text-slate-400">
Database
</span>

<span className="text-green-400">
Connected
</span>

</div>


<div className="flex justify-between">
<span className="text-slate-400">
Analytics
</span>

<span className="text-green-400">
Active
</span>

</div>


</div>


</div>

)

}
