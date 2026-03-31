import NavbarLeft from "./NavbarLeft";
import NavbarCenter from "./NavbarCenter";
import NavbarRight from "./NavbarRight";

export const Navbar = () => {
    return (
        <nav className="flex items-center gap-8 p-4 bg-gray-100">
            <NavbarLeft />

            <div className="flex-1">
                <NavbarCenter />
            </div>  

            <NavbarRight />
        </nav>
    )
}