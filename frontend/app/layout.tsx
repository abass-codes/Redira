import "./globals.css";
import Sidebar from "@/components/layout/Sidebar";

export const metadata={
title:"Redira",
description:"URL Analytics Platform"
};

export default function RootLayout({
children
}:{
children:React.ReactNode;
}){

return(
<html lang="en">
<body className="bg-black text-white">

<Sidebar/>

<main className="ml-72 min-h-screen">

<div className="w-full max-w-7xl mx-auto px-8 py-10">

{children}

</div>

</main>

</body>
</html>
);

}
