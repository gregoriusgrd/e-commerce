import { FeaturedCategories } from "./components/FeaturedCategories"
import { HeroSection } from "./components/Hero"

export const HomeView = () => {
    return (
        <div>
            <HeroSection />
            <FeaturedCategories />
        </div>
    )
}