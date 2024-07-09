import { NavLink as ReactRouterNavLink } from 'react-router-dom';

type Props = {
    path: string;
    text: string;
};

export function NavLink(
    {
        path,
        text
    }: Props
): JSX.Element {
    return (
        <ReactRouterNavLink 
            to={path}
            className={({ isActive }) => {
                return `
                    rounded-md
                    px-3
                    py-2
                    text-xl
                    font-medium
                    ${isActive ? 'bg-gray-700 text-white' : 'text-gray-300'}
                    hover:bg-gray-700
                    hover:text-white
                `;
            }}
        >
            {text}
        </ReactRouterNavLink>
    );
}