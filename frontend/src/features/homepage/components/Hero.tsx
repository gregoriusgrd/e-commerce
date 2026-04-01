// src/features/homepage/components/Hero.tsx
import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { ArrowRight } from 'lucide-react';

export function HeroSection() {
    return (
        <section className="relative w-full h-[80vh] bg-gray-100 overflow-hidden">
            <Image
                src={'/home/hero.jpg'}
                alt="hero images"
                fill
                priority
                sizes="100vw"
                className="object-cover z-0"
            />

            <div className="absolute inset-0 bg-black/20 z-10"></div>

            <div className="relative z-20 h-full flex flex-col justify-center items-center text-center px-6 max-w-5xl mx-auto">
                {/* Headline */}
                <h1 className="text-6xl md:text-8xl font-bold tracking-tight">
                    Modern Living, <br /> Refined.
                </h1>

                {/* Description */}
                <p className="mt-6 text-lg md:text-xl text-gray-200 max-w-2xl">
                    Elevating your space with timeless aesthetic and superior quality.
                </p>

                {/* Action */}
                <div className="mt-10 flex flex-col items-center gap-4">
                    <Button
                        size="lg"
                        className="cursor-pointer rounded-full px-10 py-7 text-lg bg-white text-black hover:bg-gray-200 transition-all"
                    >
                        Shop the Collection{' '}
                        <ArrowRight className="ml-2 h-5 w-5" />
                    </Button>
                </div>
            </div>
        </section>
    );
}
