import { ChangeEvent } from 'react';
import { TextFormInput } from '../TextFormInput';

type Props = {
    nameInputValue: string;
    descriptionInputValue: string;
    addressInputValue: string;
    privateKeyInputValue: string;
    onNameInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onDescriptionInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onAddressInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onPrivateKeyInputChange(event: ChangeEvent<HTMLInputElement>): void;
};

export function AddTaskForm(
    {
        nameInputValue,
        descriptionInputValue,
        addressInputValue,
        privateKeyInputValue,
        onNameInputChange,
        onDescriptionInputChange,
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
                placeholder='Name'
                value={nameInputValue}
                onChange={onNameInputChange}
            />

            <TextFormInput
                placeholder='Description'
                value={descriptionInputValue}
                onChange={onDescriptionInputChange}
            />

            <TextFormInput
                placeholder='Address'
                value={addressInputValue}
                onChange={onAddressInputChange}
            />

            <TextFormInput
                placeholder='Private Key'
                value={privateKeyInputValue}
                onChange={onPrivateKeyInputChange}
            />
        </div>
    );
}