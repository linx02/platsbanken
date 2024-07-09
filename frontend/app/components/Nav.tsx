'use client'

import Swal from "sweetalert2";
import { downloadData, downloadProgress } from "../fetchBackend";

export default function Nav() {

    function downloadDataModal() {
        Swal.fire({
            title: 'Hämta data',
            html: `
                <label for="regionInput" class="block mb-2">Hämta data efter sökterm</label>
                <input type="text" id="regionInput" class="border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
            `,
            showCancelButton: true,
            confirmButtonText: 'Hämta',
            cancelButtonText: 'Avbryt',
            cancelButtonColor: '#333333',
            confirmButtonColor: '#000056',
            focusCancel: true,
        }).then((result) => {
            if (result.isConfirmed) {
                const inputElement = document.getElementById('regionInput') as HTMLInputElement | null;
                if (inputElement === null) return; // Ensure the input element exists
                const query = inputElement.value; // Get the value of the input element
                downloadData(query); // Call downloadData with the entered query

                let interval: any; // Declare interval in a scope accessible to both didOpen and willClose

                Swal.fire({
                    title: 'Data hämtas',
                    html: 'Data hämtas från Arbetsförmedlingen...',
                    didOpen: () => {
                        Swal.showLoading();

                        interval = setInterval(async () => {
                            const progress = await downloadProgress();
                            
                            Swal.update({
                                html: `Data hämtas från Arbetsförmedlingen... ${progress}%`
                            });

                            if (progress >= 100) {
                                clearInterval(interval);
                                Swal.fire({
                                    title: 'Data hämtad',
                                    html: 'Data har hämtats från Arbetsförmedlingen!',
                                    icon: 'success',
                                });
                                localStorage.clear();
                            }
                        }, 2000); // Call downloadProgress every 2 seconds
                    },
                    willClose: () => {
                        clearInterval(interval); // Ensure the interval is cleared if the modal is closed early
                    }
                });
            }
        });
    }

    return (
        <nav className="flex justify-between margin-x items-center">
            <img src="/logo.svg" alt="Arbetsförmedlingens logga" className="" />
            <ul className="border-l-[1px] border-r-[1px] border-gray px-4 h-16 flex items-center hover:bg-[#f2f4f5] transition duration-300 hover:cursor-pointer">
                <li onClick={downloadDataModal} className="font-medium flex items-center gap-1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24">
                        <path fill="#333333" d="m12 16l-5-5l1.4-1.45l2.6 2.6V4h2v8.15l2.6-2.6L17 11zm-6 4q-.825 0-1.412-.587T4 18v-3h2v3h12v-3h2v3q0 .825-.587 1.413T18 20z"/>
                    </svg> 
                    Hämta data
                </li>
            </ul>
        </nav>
    );
}