import { Routes } from './routes';

export function App(): JSX.Element {
    return (
        <div
            className={`
                container
            `}
        >
            <Routes />
        </div>
    );
}