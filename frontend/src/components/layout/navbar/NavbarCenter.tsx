import { Search } from "lucide-react"

export default function NavbarCenter() {
    return (
        <div className="text-black flex items-center gap-2">
            <div className="relative">
                <input 
                type="text"
                placeholder="Search products..."
                className="w-full bg-white border border-gray-300 rounded-lg py-2 pl-10 pr-4 text-sm"
                />

                <Search
                size={16}
                className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
                />
            </div>

        </div>
    )
}