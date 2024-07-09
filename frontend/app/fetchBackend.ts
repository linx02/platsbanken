const url = 'http://localhost:8080';

async function handleResponse(response: Response) {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
}

export async function get(endpoint: string) {
    const response = await fetch(`${url}/${endpoint}`);
    return handleResponse(response);
}

export async function post(endpoint: string, body: any) {
    const response = await fetch(`${url}/${endpoint}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        mode: 'cors',
        body: JSON.stringify(body),
    });
    return handleResponse(response);
}

export async function downloadData(query: string) {
    try {
        const response = await post('initial-download', { query: query });
        return response.message; // Assuming the server returns an object with a 'message' property
    } catch (error) {
        console.error('Error downloading data:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}

export async function downloadProgress() {
    try {
        const response = await get('download-progress');
        return response.message; // Assuming the server returns an object with a 'message' property
    } catch (error) {
        console.error('Error fetching download progress:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}

export async function getJobCount() {
    try {
        const response = await get('job-postings');
        return response.message; // Assuming the server returns an object with a 'message' property
    } catch (error) {
        console.error('Error fetching job count:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}

export async function getJobListings(num: number) {
    try {
        return await get(`job-postings/${num}`);
    } catch (error) {
        console.error('Error fetching job listings:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}

export async function getJob(id: number) {
    try {
        return await get(`job-posting/${id}`);
    } catch (error) {
        console.error('Error fetching job:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}

export async function search(positiveSearchTerms?: Array<string>, negativeSearchTerms?: Array<string>, advancedSearchQuery?: string) {
    try {
        const body = {
            positiveSearchTerms: positiveSearchTerms,
            negativeSearchTerms: negativeSearchTerms,
            advancedSearchQuery: advancedSearchQuery,
        };

        console.log('Search body:', body);
        return await post('search', body);
    } catch (error) {
        console.error('Error performing search:', error);
        throw error; // Re-throw the error to handle it in the component
    }
}