import { ChangeEvent } from 'react';
import { TextFormInput } from './TextFormInput';

type Props = {
    nameInputValue: string;
    descriptionInputValue: string;
    emailInputValue: string;
    onNameInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onDescriptionInputChange(event: ChangeEvent<HTMLInputElement>): void;
    onEmailInputChange(event: ChangeEvent<HTMLInputElement>): void;
};

export function CreateTasklistForm(
    {
        nameInputValue,
        descriptionInputValue,
        emailInputValue,
        onNameInputChange,
        onDescriptionInputChange,
        onEmailInputChange
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
                placeholder='Email'
                value={emailInputValue}
                onChange={onEmailInputChange}
            />
        </div>
    );
}