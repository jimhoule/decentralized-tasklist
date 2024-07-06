import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { App } from './App';
import { HomePage } from './pages/HomePage';
import { TasklistsPage } from './pages/TasklistsPage';
import { WalletsPage } from './pages/WalletsPage';

const router = createBrowserRouter([
    {
        path: '/',
        element: <App />,
        children: [
            {
                index: true,
                element: <HomePage />,
            },
            {
                path: '/tasklists',
                element: <TasklistsPage />,
            },
            {
                path: '/wallets',
                element: <WalletsPage />,
            },
        ],
    },
]);

export function Router(): JSX.Element {
    return (
        <RouterProvider router={router}/>
    );
}