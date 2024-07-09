import Listing from "../elements/Listing";

export default function Results({ listings }: any) {
    return (
        <div className="py-4">
            {listings.map((job: any) => (
                <Listing
                    key={job.id}
                    title={job.title}
                    link={`/job/${job.id}`}
                    company={job.companyName}
                    location={job.location}
                    occupation={job.occupation}
                    date={job.published.split('+')[0]}
                />
            ))
            }
        </div>
    );
}