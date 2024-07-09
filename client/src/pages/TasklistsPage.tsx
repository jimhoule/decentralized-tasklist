import { useState } from 'react';
import { TaskList } from '../components/TaskList';
import { TaskListItem } from '../components/TaskListItem';
import { Modal } from '../components/Modal';
import { GetTasklistForm } from '../components/GetTasklistForm';
import { useGetTasklistForm } from '../components/GetTasklistForm/useGetTasklistForm';
import { CreateTasklistForm } from '../components/CreateTasklistForm';
import { useCreateTasklistForm } from '../components/CreateTasklistForm/useCreateTasklistForm';
import { AddTaskForm } from '../components/AddTaskForm';
import { useAddTaskForm } from '../components/AddTaskForm/useAddTaskForm';
import { Tasklist, Task, getTasklist, createTasklist, addTask } from '../api/tasklistsApi';

export function TasklistsPage(): JSX.Element {
    const { 
        getTasklistForm,
        resetGetTasklistForm,
        handleGetTasklistAddressInputChange,
        handleGetTasklistPrivateKeyInputChange
    } = useGetTasklistForm();

    const {
        createTasklistForm,
        resetCreateTasklistForm,
        handleCreateTasklistNameInputChange,
        handleCreateTasklistDescriptionInputChange,
        handleCreateTasklistEmailInputChange,
        handleCreateTasklistPrivateKeyInputChange
    } = useCreateTasklistForm();

    const {
        addTaskForm,
        resetAddTaskForm,
        handleAddTaskNameInputChange,
        handleAddTaskDescriptionInputChange,
        handleAddTaskAddressInputChange,
        handleAddTaskPrivateKeyInputChange
    } = useAddTaskForm();

    const [ tasklist, setTasklist ] = useState<Tasklist | null>(null);
    const [ isGetTasklistModalVisible, setIsGetTasklistModalVisible ] = useState<boolean>(false);
    const [ isCreateTasklistModalVisible, setIsCreateTasklistModalVisible ] = useState<boolean>(false);
    const [ isAddTaskModalVisible, setIsAddTaskModalVisible ] = useState<boolean>(false);

    return (
        <>
            <div
                className={`
                    container
                    flex
                    items-center
                    justify-center
                    bg-gray-700
                    rounded-md
                    w-1/2
                    mt-20
                    p-8
                    h-auto
                `}
            >
                <button
                    onClick={() => {
                        setIsGetTasklistModalVisible(true)
                    }}
                    className={`
                        w-auto
                        h-auto
                        bg-gray-900
                        p-2
                        rounded-md
                        text-white
                        font-bold
                        pointer-cursor
                        hover:bg-white
                        hover:text-gray-900
                    `}
                >
                    Get Tasklist
                </button>

                <button
                    onClick={() => {
                        setIsCreateTasklistModalVisible(true);
                    }}
                    className={`
                        ml-8
                        w-auto
                        h-auto
                        text-white
                        pointer-cursor
                        hover:text-gray-900
                    `}
                >
                    Dont have a Tasklist yet ? Click here to create one!
                </button>
            </div>

            {!!tasklist && (
                <TaskList 
                    tasklist={tasklist}
                    onAddTask={() => {
                        setIsAddTaskModalVisible(true)
                    }}
                    extractKey={(task: Task): string => {
                        return task.id;
                    }}
                    renderTask={(task: Task): JSX.Element => {
                        return (
                            <TaskListItem
                                key={task.id}
                                name={task.name}
                                isDone={task.isDone}
                                onUpdate={(): void => {
                                    alert('Not implemented yet!');
                                }}
                                onDelete={(): void => {
                                    alert('Not implemented yet!');
                                }}
                            />
                        );
                    }}
                />
            )}

            <Modal 
                isVisible={isGetTasklistModalVisible}
                title='Get Tasklist Form'
                submitButtonTitle='Get'
                cancelButtonTitle='Cancel'
                onSubmit={async (): Promise<void> => {
                    const tasklist = await getTasklist({
                        encodedContractAddess: getTasklistForm.address,
                        encodedPrivateKey: getTasklistForm.privateKey
                    });

                    setTasklist(tasklist);
                    setIsGetTasklistModalVisible(false);
                    resetGetTasklistForm();
                }}
                onCancel={(): void => {
                    setIsGetTasklistModalVisible(false);
                    resetGetTasklistForm();
                }}
            >
                <GetTasklistForm 
                    addressInputValue={getTasklistForm.address}
                    privateKeyInputValue={getTasklistForm.privateKey}
                    onAddressInputChange={handleGetTasklistAddressInputChange}
                    onPrivateKeyInputChange={handleGetTasklistPrivateKeyInputChange}
                />
            </Modal>

            <Modal 
                isVisible={isCreateTasklistModalVisible}
                title='Create Tasklist Form'
                submitButtonTitle='Create'
                cancelButtonTitle='Cancel'
                onSubmit={async (): Promise<void> => {
                    const tasklist = await createTasklist({
                        name: createTasklistForm.name,
                        description: createTasklistForm.description,
                        email: createTasklistForm.email,
                        encodedPrivateKey: createTasklistForm.privateKey
                    });

                    setTasklist(tasklist);
                    setIsCreateTasklistModalVisible(false);
                    resetCreateTasklistForm();
                }}
                onCancel={(): void => {
                    setIsCreateTasklistModalVisible(false);
                    resetCreateTasklistForm();
                }}
            >
                <CreateTasklistForm 
                    nameInputValue={createTasklistForm.name}
                    descriptionInputValue={createTasklistForm.description}
                    emailInputValue={createTasklistForm.email}
                    privateKeyInputValue={createTasklistForm.privateKey}
                    onNameInputChange={handleCreateTasklistNameInputChange}
                    onDescriptionInputChange={handleCreateTasklistDescriptionInputChange}
                    onEmailInputChange={handleCreateTasklistEmailInputChange}
                    onPrivateKeyInputChange={handleCreateTasklistPrivateKeyInputChange}
                />
            </Modal>

            <Modal 
                isVisible={isAddTaskModalVisible}
                title='Add Task Form'
                submitButtonTitle='Add'
                cancelButtonTitle='Cancel'
                onSubmit={async (): Promise<void> => {
                    if (!tasklist) return;

                    const task = await addTask({
                        name: addTaskForm.name,
                        description: addTaskForm.description,
                        encodedContractAddess: addTaskForm.address,
                        encodedPrivateKey: addTaskForm.privateKey
                    });

                    setTasklist({
                        ...tasklist,
                        tasks: [
                            ...(tasklist.tasks || []),
                            task
                        ]
                    });
                    setIsAddTaskModalVisible(false);
                    resetAddTaskForm();
                }}
                onCancel={(): void => {
                    setIsAddTaskModalVisible(false);
                    resetAddTaskForm();
                }}
            >
                <AddTaskForm 
                    nameInputValue={addTaskForm.name}
                    descriptionInputValue={addTaskForm.description}
                    addressInputValue={addTaskForm.address}
                    privateKeyInputValue={addTaskForm.privateKey}
                    onNameInputChange={handleAddTaskNameInputChange}
                    onDescriptionInputChange={handleAddTaskDescriptionInputChange}
                    onAddressInputChange={handleAddTaskAddressInputChange}
                    onPrivateKeyInputChange={handleAddTaskPrivateKeyInputChange}
                />
            </Modal>
        </>
    );
}