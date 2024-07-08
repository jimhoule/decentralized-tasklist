import { ReactNode } from 'react';

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
                    absolute 
                    z-1000
                    top-1/2
                    left-1/2
                    transform
                    -translate-x-1/2
                    -translate-y-1/2
                    container
                    w-screen
                    h-screen
                    bg-green-500
                    ${isVisible ? 'ease-out duration-300 opacity-100' : 'ease-in duration-200 opacity-0'}
                `}
            >
                <div className="absolute z-1000 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 px-4 pb-4 pt-5 bg-blue-600 w-1/2">
                    <div className="mt-3 text-center bg-red-400">
                        <h3
                            className={`
                                text-base
                                font-semibold
                                leading-6
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


                    <div className="bg-gray-50 px-4 py-3 bg-green">
                        <button
                            type='button'
                            className={`
                                inline-flex
                                w-1/2
                                justify-center
                                rounded-md
                                bg-red-600
                                px-3
                                py-2
                                text-sm
                                font-semibold
                                text-white
                                shadow-sm
                                hover:bg-red-500
                            `}
                            onClick={onSubmit}
                        >
                            {submitButtonTitle}
                        </button>

                        <button
                            type='button'
                            className={`
                                mt-3
                                inline-flex
                                w-1/2
                                justify-center
                                rounded-md
                                bg-white
                                px-3
                                py-2
                                text-sm
                                font-semibold
                                text-gray-900
                                shadow-sm ring-1
                                ring-inset
                                ring-gray-300
                                hover:bg-gray-50
                            `}
                            onClick={onCancel}
                        >
                            {cancelButtonTitle}
                        </button>
                    </div>
                </div>
            </div>
    );
}