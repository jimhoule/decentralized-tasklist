import { PencilSquareIcon, TrashIcon } from '@heroicons/react/24/solid';

type Props = {
    name: string;
    isDone: boolean;
    onUpdate(): void;
    onDelete(): void;
};

export function TaskListItem(
    {
        name,
        isDone,
        onUpdate,
        onDelete
    }: Props
): JSX.Element {
    return (
        <div
            className={`
                w-full
                flex
                justify-around
                items-center
                bg-violet-800
                text-white
                mb-10
                py-3
                px-4
                rounded-md
            `}
        >
            <p>{name}</p>

            <p
                className={`
                    font-bold
                    ${isDone ? 'text-green-500' : 'text-blue-300'}
                `}
            >
                {`${isDone ? 'Done' : 'In Progress'}`}
            </p>

            <button
                className={`
                    pointer-cursor
                `}
                onClick={onUpdate}
            >
                <PencilSquareIcon
                    className={`
                        size-8
                        text-blue-300
                    `}
                />
            </button>

            <button
                className={`
                    pointer-cursor
                `}
                onClick={onDelete}
            >
                <TrashIcon
                    className={`
                        size-8
                        text-red-500
                    `}
                />
            </button>
        </div>
    );
}