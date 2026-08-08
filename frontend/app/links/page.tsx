import CreateLinkForm from "@/components/forms/CreateLinkForm";
import LinkTable from "@/components/links/LinkTable";

export default function LinksPage(){

return(

<main className="min-h-screen bg-black space-y-8 p-10">

<CreateLinkForm/>

<LinkTable/>

</main>

);

}