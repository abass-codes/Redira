"use client";

import { useState } from "react";
import api from "@/lib/api";
import { useRouter } from "next/navigation";


export default function Login() {

  const router = useRouter();

  const [email,setEmail] = useState("");
  const [password,setPassword] = useState("");


  async function login(){

    const response = await api.post(
      "/auth/login",
      {
        email,
        password,
      }
    );


    localStorage.setItem(
      "token",
      response.data.token
    );


    router.push("/dashboard");

  }


  return (

    <main className="min-h-screen bg-black text-white flex items-center justify-center">

      <div className="w-96">

        <h1 className="text-3xl font-bold mb-6">
          Login
        </h1>


        <input
          className="w-full mb-3 rounded bg-gray-900 p-3"
          placeholder="Email"
          onChange={(e)=>setEmail(e.target.value)}
        />


        <input
          className="w-full mb-3 rounded bg-gray-900 p-3"
          placeholder="Password"
          type="password"
          onChange={(e)=>setPassword(e.target.value)}
        />


        <button
          onClick={login}
          className="w-full rounded bg-blue-600 p-3"
        >
          Login
        </button>

      </div>

    </main>

  );
}