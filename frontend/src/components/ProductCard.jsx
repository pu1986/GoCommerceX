"use client";
import Link from "next/link";

export default function ProductCard({ product }) {
  return (
    <div className="border rounded-lg shadow-md p-4 bg-white hover:shadow-lg transition">
      <Link href={`/product/${product._id}`}>
        <img
          src={product.image}
          alt={product.name}
          className="w-full h-48 object-cover rounded cursor-pointer"
        />
      </Link>

      <div className="mt-4">
        <Link href={`/product/${product._id}`}>
          <h2 className="text-lg font-semibold hover:underline cursor-pointer">
            {product.name}
          </h2>
        </Link>
        <p className="text-gray-600 text-sm">{product.category}</p>
        <p className="text-xl font-bold text-green-600 mt-2">₹{product.price}</p>
      </div>
    </div>
  );
}
