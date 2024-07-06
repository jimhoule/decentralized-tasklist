export function TasklistsPage(): JSX.Element {
    return (
        <div
            className={`
                container
                bg-gray-700
                mt-20
                p-8
                rounded-md
            `}
        >
            <h1
                className={`
                    text-white
                    mb-5
                `}
            >
                list name
            </h1>

            <h1
                className={`
                    text-white
                    mb-5
                `}
            >
                list description
            </h1>

            <h1
                className={`
                    text-white
                    mb-5
                `}
            >
                email
            </h1>

            <div
                className={`
                    flex
                    justify-between
                    items-center
                    bg-violet-800
                    text-white
                    py-3
                    px-4
                    rounded-md
                    mb-1
                    cursor-pointer
                `}
            >
                <p>task name</p>
                <p>is done</p>
            </div>
        </div>
    );
}