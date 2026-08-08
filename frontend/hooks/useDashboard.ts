"use client";

import {useEffect,useState} from "react";
import api from "@/lib/api";
import {DashboardSummary} from "@/types/dashboard";

export default function useDashboard(){

const [dashboard,setDashboard]=useState<DashboardSummary|null>(null);

const [loading,setLoading]=useState(true);


useEffect(()=>{

async function load(){

try{

const response=await api.get("/dashboard");

setDashboard(response.data);

}finally{

setLoading(false);

}

}

load();

},[]);


return{
dashboard,
loading
};

}