type ListingProps = {
    title: string;
    link: string;
    company: string;
    location: string;
    occupation: string;
    date: string;
};

export default function Listing({ title, link, company, location, occupation, date }: ListingProps) {
    return (
        <div className="block bg-white p-4 margin-x !my-3">
            <a href={link} className="text-xl text-[#1616ab] underline font-semibold">{title}</a>
            <p className="text-md font-medium py-2">{company} - {location}</p>
            <p className="text-md">{occupation}</p>
            <p className="text-md">Publicerad: {date}</p>
        </div>
    );
}