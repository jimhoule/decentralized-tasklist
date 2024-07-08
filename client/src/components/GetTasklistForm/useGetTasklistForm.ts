import { useState, ChangeEvent } from 'react';

type GetTasklistForm = {
    address: string;
    privateKey: string;
};

export function useGetTasklistForm() {
    const [ getTasklistForm, setGetTasklistForm ] = useState<GetTasklistForm>({
        address: '',
        privateKey: ''
    });

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
        handleGetTasklistAddressInputChange,
        handleGetTasklistPrivateKeyInputChange,
    };
}