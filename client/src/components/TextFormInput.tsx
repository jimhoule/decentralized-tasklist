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
            placeholder={placeholder}
            value={value}
            onChange={onChange}
        />
    );
}