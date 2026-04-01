import Image from 'next/image';
import Link from 'next/link';

export function FeaturedCategories() {
    return (
        <section className="mx-auto max-w-6xl px-4 py-12 text-black">
            <div className="flex items-center gap-8">
                <div className="group relative w-full max-w-[321px] h-[584px] bg-gray-50 overflow-hidden rounded-xl border border-gray-100 transition-all hover:shadow-xl">
                    {/* Image Container - Mengisi sekitar 75-80% tinggi box */}
                    <div className="relative flex-1 bg-[#F5F5F5] overflow-hidden">
                        <Image
                            src="/home/categories_1.jpg"
                            alt="Chair Category"
                            fill
                            className="object-contain p-8 transition-transform duration-500 group-hover:scale-110"
                        />
                    </div>

                    {/* Text Content */}
                    <div className="flex flex-col justify-between p-6 h-[30%] bg-white">
                        <div>
                            <h3 className="text-xl font-bold text-gray-900">
                                Premium Chairs
                            </h3>
                            <p className="text-sm text-gray-500 mt-1">
                                Comfortable and stylish seating for your modern
                                home.
                            </p>
                        </div>

                        <Link
                            href="#"
                            className="text-sm font-semibold underline underline-offset-4 hover:text-gray-600 flex items-center gap-2"
                        >
                            View More
                        </Link>
                    </div>
                </div>


                <div className="group w-full max-w-[321px] h-[584px] flex flex-col overflow-hidden rounded-xl border border-gray-100 shadow-sm transition-hover hover:shadow-md">
                    {/* 1. Image */}
                    <div className="relative flex-1 bg-[#F5F5F5] overflow-hidden">
                        <Image
                            src="/home/categories_2.jpg"
                            alt="Sofa"
                            fill
                        />
                    </div>

                    {/* 2. Text */}
                    <div className="bg-white p-6 flex flex-col gap-2">
                        <h3 className="text-xl font-bold text-gray-900">
                            Elegant Sofas
                        </h3>
                        <p className="text-sm text-gray-500 line-clamp-2">
                            Luxurious sofas that combine comfort with timeless
                            design.
                        </p>
                        <Link
                            href="#"
                            className="mt-2 text-sm font-semibold underline underline-offset-4 hover:text-black"
                        >
                            View More
                        </Link>
                    </div>
                </div>

                <div className="group relative w-full max-w-[321px] h-[584px] bg-gray-50 overflow-hidden rounded-xl border border-gray-100 transition-all hover:shadow-xl">
                    <div className="relative w-full h-[70%] overflow-hidden bg-[#F5F5F5]">
                        <Image
                            src="/home/categories_3.jpg"
                            fill
                            alt="Table"
                            className="object-contain p-8 transition-transform duration-500 group-hover:scale-110"
                        />
                    </div>

                    <div className="flex flex-col justify-between p-6 h-[30%] bg-white">
                        <div>
                            <h3 className="text-xl font-bold text-gray-900">
                                Modern Tables
                            </h3>
                            <p className="text-sm text-gray-500 mt-1">
                                Stylish tables that blend functionality with
                                contemporary design.
                            </p>
                        </div>
                        <Link
                            href="#"
                            className="text-sm font-semibold underline underline-offset-4 hover:text-gray-600 flex items-center gap-2"
                        >
                            View More
                        </Link>
                    </div>
                </div>
            </div>
        </section>
    );
}
