import { useState, ChangeEvent } from 'react';

type AddTaskForm = {
    name: string;
    description: string;
    address: string;
    privateKey: string;
};

const initialAddTaskForm = {
    name: '',
    description: '',
    address: '',
    privateKey: ''
}

export function useAddTaskForm() {
    const [ addTaskForm, setAddTaskForm ] = useState<AddTaskForm>(initialAddTaskForm);

    function resetAddTaskForm(): void {
        setAddTaskForm(initialAddTaskForm);
    }

    function handleAddTaskNameInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setAddTaskForm({
            ...addTaskForm,
            name: event.currentTarget.value
        });
    }

    function handleAddTaskDescriptionInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setAddTaskForm({
            ...addTaskForm,
            description: event.currentTarget.value
        });
    }

    function handleAddTaskAddressInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setAddTaskForm({
            ...addTaskForm,
            address: event.currentTarget.value
        });
    }

    function handleAddTaskPrivateKeyInputChange(event: ChangeEvent<HTMLInputElement>): void {
        setAddTaskForm({
            ...addTaskForm,
            privateKey: event.currentTarget.value
        });
    }

    return {
        addTaskForm,
        resetAddTaskForm,
        handleAddTaskNameInputChange,
        handleAddTaskDescriptionInputChange,
        handleAddTaskAddressInputChange,
        handleAddTaskPrivateKeyInputChange
    };
}