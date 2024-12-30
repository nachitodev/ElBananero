import { Menu, Search } from "lucide-react";

function Navbar(): JSX.Element {
    return (
        <div className="flex justify-center bg-black h-12">
            <img src="/img/logo_bananero.png" className="" />
            <Search color="#ffffff" strokeWidth={2.25} className="m-4 p-1" />
            <Menu color="#ffffff" className="m-4 p-1 bg-red-600 hover:bg-red-700 cursor-pointer" />
        </div>
    )
}

export default Navbar;