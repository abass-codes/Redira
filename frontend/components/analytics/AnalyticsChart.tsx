export default function AnalyticsChart(){

  return(
  
  <div className="rounded-2xl border border-slate-800 bg-slate-950 p-8">
  
  <h2 className="text-2xl font-bold text-white">
  Analytics Overview
  </h2>
  
  <p className="mt-2 text-slate-400">
  Track performance trends.
  </p>
  
  <div className="mt-8 flex h-56 items-end gap-5 rounded-xl bg-slate-900 p-6">
  
  <div className="h-20 w-full rounded bg-blue-600"/>
  
  <div className="h-32 w-full rounded bg-blue-500"/>
  
  <div className="h-44 w-full rounded bg-purple-500"/>
  
  <div className="h-28 w-full rounded bg-pink-500"/>
  
  </div>
  
  </div>
  
  );
  
  }