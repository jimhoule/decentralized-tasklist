import { ChangeEvent } from 'react';

type Props = {
    placeholder: string;
    value: string;
    onChange(event: ChangeEvent<HTMLInputElement>): void
};

export function TextFormInput(
    {
        placeholder,
        value,
        onChange
    }: Props
): JSX.Element {
    return (
        <input
            type='text'
            className={`
                mb-2
                px-1
                py-2
                border-2
                border-gray-900
                rounded-md
                w-3/4
                focus:placeholder-white
                focus:bg-gray-400
                focus:text-white
                focus:py-3
                ease-in-out
                duration-200
            `}
            placeholder={placeholder}
            value={value}
            onChange={onChange}
        />
    );
}