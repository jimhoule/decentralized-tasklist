import { axiosInstance } from './axios';

export type Task =  {
    id: string,
    name: string,
    description?: string,
    isDone: boolean
};

export type Tasklist = {
    name: string,
    description: string,
    email: string,
    tasks: Task[]
};

export type GetTasklistPayload = {
    encodedContractAddess: string;
    encodedPrivateKey: string;
};
export type CreateTasklistPayload = Omit<Tasklist, 'tasks'> & Pick<GetTasklistPayload, 'encodedPrivateKey'>;
export type AddTaskPayload = Required<Pick<Task, 'name' | 'description'>> & GetTasklistPayload;
export type UpdateTaskPayload = Required<Task> & GetTasklistPayload;
export type DeleteTaskPayload = Pick<Task, 'id'> & GetTasklistPayload;

const endpoint = '/tasklists';

export async function getTasklist(
    {
        encodedContractAddess,
        encodedPrivateKey
    }: GetTasklistPayload
): Promise<Tasklist> {
    const { data } = await axiosInstance.get<Tasklist>(`${endpoint}/${encodedContractAddess}/${encodedPrivateKey}`);

    return data;
}

export async function createTasklist(
    {
        name,
        description,
        email,
        encodedPrivateKey
    }: CreateTasklistPayload
): Promise<Tasklist> {
    const { data } = await axiosInstance.post<Tasklist>(`${endpoint}/${encodedPrivateKey}`, {
        name,
        description,
        email
    });

    return data;
}

export async function addTask(
    {
        name,
        description,
        encodedContractAddess,
        encodedPrivateKey
    }: AddTaskPayload
): Promise<Task> {
    const { data } = await axiosInstance.post<Task>(`${endpoint}/tasks/${encodedContractAddess}/${encodedPrivateKey}`, {
        name,
        description,
    });

    return data;
}

export async function updateTask(
    {
        id,
        name,
        description,
        isDone,
        encodedContractAddess,
        encodedPrivateKey
    }: UpdateTaskPayload
): Promise<Task> {
    const { data } = await axiosInstance.put<Task>(`${endpoint}/tasks/${id}/${encodedContractAddess}/${encodedPrivateKey}`, {
        name,
        description,
        isDone
    });

    return data;
}

export async function deleteTask(
    {
        id,
        encodedContractAddess,
        encodedPrivateKey
    }: DeleteTaskPayload
): Promise<void> {
    await axiosInstance.delete<string>(`${endpoint}/tasks/${id}/${encodedContractAddess}/${encodedPrivateKey}`);
}