interface StatsCardProps {
  title: string;
  value: number;
  description: string;
}


export default function StatsCard({
  title,
  value,
  description,
}: StatsCardProps){

return(

<div className="
rounded-2xl
border
border-slate-800
bg-slate-950
p-6
shadow-lg
">

<p className="
text-sm
text-slate-400
">
{title}
</p>


<p className="
mt-3
text-4xl
font-bold
text-white
">
{value}
</p>


<p className="
mt-2
text-sm
text-slate-500
">
{description}
</p>


</div>

);

}