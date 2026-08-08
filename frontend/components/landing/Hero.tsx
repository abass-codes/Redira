import UrlShortenerDemo from "./UrlShortenerDemo";

export default function Hero(){

return(

<section className="flex min-h-screen items-center justify-center px-6">

<div className="max-w-5xl text-center">


<p className="mb-6 text-sm uppercase tracking-[0.35em] text-blue-400">
Developer Infrastructure Platform
</p>


<h1 className="text-6xl font-bold leading-tight text-white md:text-8xl">

Shorten links.

<br/>

<span className="bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">
Track performance.
</span>

</h1>


<p className="mx-auto mt-8 max-w-3xl text-xl text-slate-300">

A production-ready URL shortening platform with analytics,
performance tracking, and developer APIs.

</p>


<div className="mt-12">

<UrlShortenerDemo/>

</div>


</div>

</section>

);

}