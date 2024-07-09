import { Tasklist, Task } from '../api/tasklistsApi';

type Props = {
    tasklist: Tasklist;
    onAddTask(): void;
    extractKey(task: Task, index: number): string;
    renderTask(task: Task, index: number): JSX.Element;
};

export function TaskList(
    {
        tasklist,
        onAddTask,
        extractKey,
        renderTask
    }: Props
): JSX.Element {
    return (
        <div
            className={`
                container
                bg-gray-700
                mt-20
                p-8
                rounded-md
                h-[60%]
            `}
        >
            <div
                className={`
                    flex
                    items-center
                    justify-around
                    py-3
                    text-2xl
                    text-white
                `}
            >
                <p>Name: {tasklist.name}</p>
                <p>Email: {tasklist.email}</p>
            </div>

            <div
                className={`
                    mt-3
                    flex
                    items-center
                    justify-start
                    p-3
                    text-xl
                    text-white
                `}
            >
                <p>{tasklist.description}</p>
            </div>

            <button
                onClick={onAddTask}
                className={`
                    w-auto
                    h-auto
                    bg-gray-900
                    p-2
                    mt-10
                    rounded-md
                    text-white
                    font-bold
                    pointer-cursor
                    hover:bg-white
                    hover:text-gray-900
                `}
            >
                Add Task
            </button>

            <div
                className={`
                    container
                    mt-3
                    p-4
                    text-xl
                    text-white
                    border-gray-900
                    border-4
                    rounded-md
                    h-[55%]
                    overflow-y-auto
                `}
            >
                {tasklist?.tasks?.map((task: Task, index: number): JSX.Element => {
                    return (
                        <div
                            key={extractKey(task, index)}
                        >
                            {renderTask(task, index)}
                        </div>
                    );
                })}
            </div>
        </div>
    );
}