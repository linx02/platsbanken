'use client'

import Search from "./components/blocks/Search";
import Overview from "./components/blocks/Overview";
import Results from "./components/blocks/Results";
import { getJobListings, search } from "./fetchBackend";
import { useState, useEffect } from "react";

export default function Home() {

  const [listings, setListings] = useState([]);

  useEffect(() => {
    const fetchListings = async () => {
      // Check if listings are already stored in localStorage
      const storedListings = localStorage.getItem("listings");
      if (storedListings) {
        setListings(JSON.parse(storedListings));
      } else {
        const fetchedListings = await getJobListings(25);
        setListings(fetchedListings);
        // Save fetched listings to localStorage
        localStorage.setItem("listings", JSON.stringify(fetchedListings));
      }
    };

    fetchListings();
  }, []);

  const handleSearch = async (positiveSearchTerms: string[], negativeSearchTerms: string[], advancedSearchQuery: string) => {
    try {
      const searchedListings = await search(positiveSearchTerms, negativeSearchTerms, advancedSearchQuery);
      setListings(searchedListings);
      // Save searched listings to localStorage
      localStorage.setItem("listings", JSON.stringify(searchedListings));
    } catch (error) {
      console.error("Error performing search:", error);
    }
  };

  return (
    <main className="bg-light-gray">
      <Search onSearch={handleSearch} />
      <Overview listingCount={listings ? listings.length : -1} />
      { listings && <Results listings={listings} /> }
    </main>
  );
}
