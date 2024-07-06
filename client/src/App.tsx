import { Outlet } from 'react-router-dom';
import { NavBar } from './components/NavBar';

export function App(): JSX.Element {
    return (
        <div
            className={`
                container
                w-screen
                h-screen
                max-w-full
                flex
                flex-col
            `}
        >
            <div
                className={`
                    container
                    w-full
                    h-auto
                    max-w-full
                    bg-black
                `}
            >
                <NavBar />
            </div>

            <div
                className={`
                    container
                    w-full
                    h-full
                    max-w-full
                    bg-black
                `}
            >
                <Outlet/>
            </div>
        </div>
    );
}