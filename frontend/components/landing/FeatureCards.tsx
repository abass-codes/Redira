const features=[

    {
    title:"Fast",
    description:"Instant redirects with optimized backend performance."
    },
    
    {
    title:"Analytics",
    description:"Track clicks, devices, locations, and traffic sources."
    },
    
    {
    title:"Reliable",
    description:"Built with production-grade infrastructure."
    }
    
    ];
    
    
    export default function FeatureCards(){
    
    return(
    
    <section className="px-6 py-20">
    
    
    <div className="mx-auto grid max-w-5xl gap-6 md:grid-cols-3">
    
    
    {features.map((feature)=>(
    
    
    <div
    key={feature.title}
    className="rounded-2xl border border-slate-800 bg-slate-950 p-8 text-center shadow-lg"
    >
    
    
    <h2 className="text-2xl font-bold text-white">
    {feature.title}
    </h2>
    
    
    <p className="mt-4 text-slate-400">
    {feature.description}
    </p>
    
    
    </div>
    
    
    ))}
    
    
    </div>
    
    
    </section>
    
    );
    
    }