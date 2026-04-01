import { ProductCard } from "@/src/features/product/components/ProductCard";
import { ProductList } from "@/src/features/product/components/ProductList";
import { HomeView } from "@/src/features/homepage";

export default function Home() {
    const mockProduct = [{
        id: 1,
        name: "White Chair",
        price: 500000,
        image: "/temp/white chair.jpg"
    },
    {
        id: 2,
        name: "Modern Sofa",
        price: 1500000,
        image: "/temp/red sofa.jpg"
    },
    {
        id: 3,
        name: "Hanger",
        price: 30000,
        image: "/temp/hanger.jpg"
    },
    {
        id: 4,
        name: "Wardrobe",
        price: 2000000,
        image: "/temp/wardrobe.jpg"
    },

    ];

    return (
        <main className=" bg-white min-h-screen">
            <HomeView />
            <h1 className="text-2xl font-bold mb-4">Products</h1>
            
            <ProductList products={mockProduct} />
        </main>
    )
}