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
            <svg width="250" height="80" viewBox="0 0 250 80" xmlns="http://w3.org">
                <text x="10" y="55" fontFamily="Arial, sans-serif" fontSize="40" fontWeight="bold" fill="black">
                    ALIT
                </text>

                <path d="M140 45 C 140 30, 160 30, 160 45 C 160 60, 140 70, 140 75 C 140 70, 120 60, 120 45 C 120 30, 140 30, 140 45 Z" fill="#800080" />

                <path d="M190 45 C 190 30, 210 30, 210 45 C 210 60, 190 70, 190 75 C 190 70, 170 60, 170 45 C 170 30, 190 30, 190 45 Z" fill="#800080" />
            </svg>

            <NavbarRight />
        </nav>
    )
}