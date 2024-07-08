import { useState, ChangeEvent } from 'react';

type CreateTasklistForm = {
    name: string;
    description: string;
    email: string;
};

export function useCreateTasklistForm() {
    const [ createTasklistForm, setCreateTasklistForm ] = useState<CreateTasklistForm>({
        name: '',
        description: '',
        email: ''
    });

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

    return {
        createTasklistForm,
        handleCreateTasklistNameInputChange,
        handleCreateTasklistDescriptionInputChange,
        handleCreateTasklistEmailInputChange
    };
}