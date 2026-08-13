import LoginForm from "@/components/forms/LoginForm";

export default function Login(){

return(

<main className="relative flex min-h-screen items-center justify-center overflow-hidden bg-slate-50 px-6">

<div className="absolute inset-0 bg-[radial-gradient(circle_at_top,_#dbeafe,_transparent_40%)]" />

<div className="relative z-10 flex w-full flex-col items-center">

<LoginForm />

<footer className="mt-8 text-center text-xs text-slate-400">
© 2026 Redira. Built by Yakubu Mohammed Abass.
</footer>

</div>

</main>

);

}
