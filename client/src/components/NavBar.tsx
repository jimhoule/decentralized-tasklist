import { NavLink } from './NavLink';
import ethPng from '../assets/eth.png';

export function NavBar(): JSX.Element {
    return (
        <nav
            className={`
                bg-gray-900
                w-full
                h-full
                p-2
                flex
            `}
        >
            <div
                className={`
                    flex
                    items-center
                    mr-10
                `}
            >
                <img 
                    className={`
                        h-12 
                        w-auto
                    `}
                    src={ethPng}
                    alt='eth'
                />
            </div>

            <div
                className={`
                    flex
                    space-x-4
                `}
            >
                <NavLink path='/' text='Home' />
                <NavLink path='/tasklists' text='Tasklists' />
                <NavLink path='/wallets' text='Wallets' />
            </div>
        </nav>
    );
}