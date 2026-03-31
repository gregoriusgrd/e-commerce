import { Heart, ShoppingCart } from 'lucide-react';
import Link from 'next/link';

export default function NavbarRight() {
    return (
        <div className="text-black flex items-center gap-4">
            <Link 
            href="/cart" 
            className="flex items-center gap-2 hover:text-gray-600 transition-colors">
                <ShoppingCart size={18} />
                <span className='text-sm uppercase'>cart</span>
            </Link>
            <Link 
            href="/wishlist" 
            className="flex items-center gap-2 hover:text-gray-600 transition-colors">
                <Heart size={18} />
                <span className='text-sm uppercase'>wishlist</span>
            </Link>
        </div>
    );
}
