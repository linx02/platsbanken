'use client'

import Input from "../elements/Input";

export default function Search({ onSearch }: any) {

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.target as HTMLFormElement);
        const entries = Object.fromEntries(formData.entries());
    
        // Create a new object with the correct types
        const data: { [key: string]: any } = {};
    
        // Copy entries to data object, split specific fields, and remove empty values
        for (const [key, value] of Object.entries(entries)) {
            if (key === 'positiveSearchTerms' || key === 'negativeSearchTerms') {
                // Transform specific string values by splitting at commas and remove empty terms
                if (typeof value === 'string') {
                    const splitValues = value.split(',').map(term => term.trim()).filter(term => term);
                    if (splitValues.length > 0) {
                        data[key] = splitValues;
                    }
                }
            } else {
                // Add non-empty values to the data object
                if (value !== '') {
                    data[key] = value;
                }
            }
        }
    
        // Send the data to the backend
        onSearch(data.positiveSearchTerms, data.negativeSearchTerms, data.advancedSearchQuery);

        console.log(data);
    };

    return (
        <div className="bg-dark-blue padding-x !py-12">
            <h1 className="text-white text-3xl font-semibold">Platsbanken - fast bättre</h1>
            <div className="pt-12">
                <p className="text-white text-md font-semibold">Sök på ett eller flera ord</p>
                <p className="text-white text-md">Skriv t.ex. [utvecklare], [flerårig erfarenhet], [not ".net" in description]</p>
            </div>
            <form className="flex gap-4 pt-12 w-full items-end" onSubmit={handleSubmit}>
                <Input type="text" name="positiveSearchTerms" label="Positiv sökning" value="" />
                <Input type="text" name="negativeSearchTerms" label="Negativ sökning" value="" />
                <Input type="text" name="advancedSearchQuery" label="Avancerad sökning (experimentell)" value="" />
                <button type="submit" className="flex items-center gap-1 bg-blue text-dark-blue font-semibold px-4 py-2 rounded-md bg-[#a7e95f] h-12"><svg xmlns="http://www.w3.org/2000/svg" width="1.2em" height="1.2em" viewBox="0 0 12 12"><path fill="#000056" d="M5 1a4 4 0 1 0 2.248 7.31l2.472 2.47a.75.75 0 1 0 1.06-1.06L8.31 7.248A4 4 0 0 0 5 1M2.5 5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0"/></svg> Sök</button>
            </form>
        </div>
    );
}