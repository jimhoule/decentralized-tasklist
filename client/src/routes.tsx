import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { HomePage } from './pages/HomePage';
import { TasklistsPage } from './pages/TasklistsPage';
import { WalletsPage } from './pages/WalletsPage';

const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />,
    },
    {
        path: "/tasklists",
        element: <TasklistsPage />,
    },
    {
        path: "/wallets",
        element: <WalletsPage />,
    },
]);

export function Routes(): JSX.Element {
    return (
        <RouterProvider router={router}/>
    );
}