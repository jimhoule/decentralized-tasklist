import { ReactNode, Suspense } from 'react';

type Props = {
    children: ReactNode;
    isVisible: boolean;
    title: string;
    submitButtonTitle: string;
    cancelButtonTitle: string;
    onSubmit(): void,
    onCancel(): void;
};

export function Modal(
    {
        children,
        isVisible,
        title,
        submitButtonTitle,
        cancelButtonTitle,
        onSubmit,
        onCancel
    }: Props
): JSX.Element {
    return (
        <div 
            className={`
                container
                absolute 
                z-1000
                top-1/2
                left-1/2
                transform
                -translate-x-1/2
                -translate-y-1/2
                w-1/2
                h-screen
                bg-gray-900
                opacity-90
                ${isVisible ? 'ease-out duration-300 top-1/2' : 'ease-in duration-200 top-[-100%]'}
            `}
        >
            <Suspense
                    fallback={<div>Loading....</div>}
            >
                <div 
                    className={`
                        absolute
                        z-1000
                        top-1/2
                        left-1/2
                        transform
                        -translate-x-1/2
                        -translate-y-1/2
                        px-4
                        pb-4
                        pt-5
                        bg-white
                        w-1/2"
                    `}
                >
                    <div
                        className={`
                            mt-3
                            text-center
                            bg-white
                        `}
                    >
                        <h3
                            className={`
                                font-bold
                                text-gray-900
                            `}
                        >
                            {title}
                        </h3>

                        <div 
                            className={`
                                mt-2
                            `}
                        >
                            {children}
                        </div>
                    </div>


                    <div
                        className={`
                            mt-5
                            flex
                            justify-around
                        `}
                    >
                        <button
                            className={`
                                w-[40%]
                                rounded-md
                                bg-gray-900
                                py-2
                                font-semibold
                                text-white
                                hover:bg-gray-600
                                focus:bg-gray-600
                            `}
                            onClick={onSubmit}
                        >
                            {submitButtonTitle}
                        </button>

                        <button
                            className={`
                                w-[40%]
                                rounded-md
                                bg-gray-900
                                py-2
                                font-semibold
                                text-white
                                hover:bg-gray-600
                                focus:bg-gray-600
                            `}
                            onClick={onCancel}
                        >
                            {cancelButtonTitle}
                        </button>
                    </div>
                </div>
            </Suspense>
        </div>
    );
}