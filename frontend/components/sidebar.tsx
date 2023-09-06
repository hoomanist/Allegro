import {FiHome} from "@react-icons/all-files/fi/FiHome";
import {FiFeather} from "@react-icons/all-files/fi/FiFeather";
import Link from "next/link";

export default function SideBar() {
    return (
        <div className="fixed top-0 left-0 h-screen w-16 m-0 flex flex-col
        bg-zinc-800 text-white shadow">
            <Link href="/">
                <div className='sidebar-icon rounded-lg group'>
                    <FiHome />
                    <span className="sidebar-tooltip group-hover:scale-100"> Home </span>
                </div>
            </Link>


            <Link href="/composers">
                <div className='sidebar-icon rounded-full scale-[0.8] group'>
                    <FiFeather />
                    <span className="sidebar-tooltip group-hover:scale-100"> Composers </span>
                </div>
            </Link>
        </div>
    )
}
