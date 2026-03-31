import { Product } from "../types";
import { ProductCard } from "./ProductCard";

interface ProductListProps {
    products: Product[];
}

export const ProductList = ({ products }: ProductListProps) => {
    return (
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            {products.map((product) => (
                <ProductCard key={product.id} product={product} />
            ))}
        </div>
    )
}