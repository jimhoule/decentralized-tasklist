import { useState, ChangeEvent } from 'react';

type GetTasklistForm = {
    address: string;
    privateKey: string;
};

const initialGetTasklistFrom = {
    address: '',
    privateKey: ''
}

export function useGetTasklistForm() {
    const [ getTasklistForm, setGetTasklistForm ] = useState<GetTasklistForm>(initialGetTasklistFrom);

    function resetGetTasklistForm(): void {
        setGetTasklistForm(initialGetTasklistFrom);
    }

    function handleGetTasklistAddressInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setGetTasklistForm({
            ...getTasklistForm,
            address: event.currentTarget.value
        });
    }

    function handleGetTasklistPrivateKeyInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setGetTasklistForm({
            ...getTasklistForm,
            privateKey: event.currentTarget.value
        });
    }

    return {
        getTasklistForm,
        resetGetTasklistForm,
        handleGetTasklistAddressInputChange,
        handleGetTasklistPrivateKeyInputChange,
    };
}