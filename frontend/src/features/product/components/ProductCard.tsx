// src/features/product/components/ProductCard.tsx
import { Product } from "../types";

interface ProductCardProps {
  product: Product;
}

export const ProductCard = ({ product }: ProductCardProps) => {
  return (
    <div className="border rounded-xl p-4 shadow-sm hover:shadow-md transition">
      <img 
        src={product.image} 
        alt={product.name}
        className="w-full h-40 object-cover rounded-md"
        />

        <div className="mt-3 space-y-2">
          <h3 className="text-sm font-semibold">{product.name}</h3>

            <p className="text-lg font-bold text-blue-600">
              Rp {product.price.toLocaleString()}
            </p>

            <button className="w-full bg-black text-white py-2 rounded-md hover:bg-gray-800">
              Add to Cart
            </button>
        </div>
    </div>
  );
};