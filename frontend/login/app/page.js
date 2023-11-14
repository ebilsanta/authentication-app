import Hero from './components/hero'
import Menu from './components/menu'

export default function Home() {
  return <div className="flex sm:flex-row flex-col w-screen h-screen">
    <Hero />
    <Menu />
  </div>
}
