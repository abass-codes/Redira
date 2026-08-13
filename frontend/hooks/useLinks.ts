"use client";

import {useEffect,useState} from "react";
import api from "@/lib/api";

export interface Link{

ID:string;

OriginalUrl:string;

ShortCode:string;

ClickCount:number;

}


export default function useLinks(){

const [links,setLinks]=useState<Link[]>([]);

const [loading,setLoading]=useState(true);


async function loadLinks(){

try{

const response=await api.get("/links");

setLinks(response.data.links ?? []);

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