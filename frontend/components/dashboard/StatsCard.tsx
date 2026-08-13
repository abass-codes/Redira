export default function StatsCard({
title,
value,
subtitle
}:{
title:string;
value:number;
subtitle:string;
}){

return(

<div className="
rounded-2xl
border
border-white/10
bg-gradient-to-br
from-slate-900
to-slate-950
p-6
shadow-xl
">

<p className="text-sm text-slate-400">
{title}
</p>

<h2 className="
mt-3
text-4xl
font-bold
text-white
">
{value}
</h2>

<p className="
mt-2
text-sm
text-slate-500
">
{subtitle}
</p>

</div>

)

}
