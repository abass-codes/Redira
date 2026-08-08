export default function Input(
    props:React.InputHTMLAttributes<HTMLInputElement>
    ){
    
    return(
    
    <input
    {...props}
    className="rounded-xl border border-slate-300 bg-white px-5 py-3 text-slate-900 outline-none focus:border-blue-500"
    />
    
    );
    
    }