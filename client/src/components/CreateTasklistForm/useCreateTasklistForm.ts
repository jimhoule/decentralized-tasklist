import { useState, ChangeEvent } from 'react';

type CreateTasklistForm = {
    name: string;
    description: string;
    email: string;
    privateKey: string;
};

const initialCreateTasklistForm = {
    name: '',
    description: '',
    email: '',
    privateKey: ''
}

export function useCreateTasklistForm() {
    const [ createTasklistForm, setCreateTasklistForm ] = useState<CreateTasklistForm>(initialCreateTasklistForm);

    function resetCreateTasklistForm(): void {
        setCreateTasklistForm(initialCreateTasklistForm);
    }

    function handleCreateTasklistNameInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setCreateTasklistForm({
            ...createTasklistForm,
            name: event.currentTarget.value
        });
    }

    function handleCreateTasklistDescriptionInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setCreateTasklistForm({
            ...createTasklistForm,
            description: event.currentTarget.value
        });
    }

    function handleCreateTasklistEmailInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setCreateTasklistForm({
            ...createTasklistForm,
            email: event.currentTarget.value
        });
    }

    function handleCreateTasklistPrivateKeyInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setCreateTasklistForm({
            ...createTasklistForm,
            privateKey: event.currentTarget.value
        });
    }

    return {
        createTasklistForm,
        resetCreateTasklistForm,
        handleCreateTasklistNameInputChange,
        handleCreateTasklistDescriptionInputChange,
        handleCreateTasklistEmailInputChange,
        handleCreateTasklistPrivateKeyInputChange
    };
}