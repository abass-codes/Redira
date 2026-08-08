"use client";

import {useEffect,useState} from "react";
import api from "@/lib/api";

export interface Link{

id:string;

original_url:string;

short_code:string;

clicks:number;

}


export default function useLinks(){

const [links,setLinks]=useState<Link[]>([]);

const [loading,setLoading]=useState(true);


async function loadLinks(){

try{

const response=await api.get("/links");

setLinks(response.data);

}finally{

setLoading(false);

}

}


useEffect(()=>{

loadLinks();

},[]);


return{

links,
loading,
refresh:loadLinks

};

}