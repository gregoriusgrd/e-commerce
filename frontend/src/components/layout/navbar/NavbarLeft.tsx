import Image from 'next/image';

export default function NavbarLeft() {
    return (
        <div className="flex items-center gap-4">
            <Image
                src="/logo/livingroom.png"
                alt="Logo"
                width={50}
                height={50}
            />
            <div>
                <h2 className="text-xl font-bold text-black">Furniture</h2>
                <p className="text-sm text-gray-900">Your Future Interior</p>
            </div>
        </div>
    );
}
