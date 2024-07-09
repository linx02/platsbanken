'use client'

import { useState, useEffect } from "react";

type InputProps = {
    type: "text" | "password" | "email";
    label: string;
    value: string;
    name: string;
};

export default function Input({ type, label, value, name }: InputProps) {

    const [inputValue, setInputValue] = useState(value);

    useEffect(() => {
        setInputValue(value);
    }, [value]);

    return (
        <div className="flex flex-col w-full">
            <label className="text-sm font-semibold text-white pb-2">{label}</label>
            <input
                type={type}
                value={inputValue}
                name={name}
                onChange={(event) => setInputValue(event.target.value)}
                className="border border-light-gray rounded-md p-2 h-12"
            />
        </div>
    );
}