import { ChangeEvent } from 'react';
import { TextFormInput } from '../TextFormInput';

type Props = {
    addressInputValue: string;
    privateKeyInputValue: string;
    onAddressInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onPrivateKeyInputChange(event: ChangeEvent<HTMLInputElement>): void;
};

export function GetTasklistForm(
    {
        addressInputValue,
        privateKeyInputValue,
        onAddressInputChange,
        onPrivateKeyInputChange
    }: Props
): JSX.Element {
    return (
        <div
            className={`
                container
            `}
        >
            <TextFormInput
                placeholder='Encoded contract address'
                value={addressInputValue}
                onChange={onAddressInputChange}
            />

            <TextFormInput
                placeholder='Encoded private key'
                value={privateKeyInputValue}
                onChange={onPrivateKeyInputChange}
            />
        </div>
    );
}