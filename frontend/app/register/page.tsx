import RegisterForm from "@/components/forms/RegisterForm";

export default function Register(){

return(

<main className="min-h-screen bg-slate-50 flex items-center justify-center px-6">

<div className="w-full max-w-md rounded-3xl border border-slate-200 bg-white p-10 shadow-xl">

<h1 className="text-5xl font-bold text-slate-900">
Create Account
</h1>

<RegisterForm />

</div>

</main>

);

}