'use client'

import { getJobCount } from "@/app/fetchBackend";
import { useEffect, useState } from "react";

export default function Overview({ listingCount }: { listingCount: number }) {
    const [jobCount, setJobCount] = useState(0);

    useEffect(() => {
        const fetchJobCount = async () => {
            const count = await getJobCount();
            setJobCount(count);
        }

        fetchJobCount();
    }, []);

    return (
        <div className="padding-x !py-8 bg-white">
            { jobCount > 0 ? <h2 className="text-dark-blue text-2xl font-bold">{listingCount} av {jobCount} annonser</h2> :
            <em className="text-dark-blue text-2xl font-bold">Databasen är tom, börja med att hämta data</em>
            }
        </div>
    )
}