import SvgIcon from '../components/SvgIcon'

const Logo: React.FC<React.HTMLAttributes<HTMLDivElement>> = () => {
  return (
    <header className="flex items-center px-6 h-20 bg-gray-dark border-r border-color-base">
      <SvgIcon name="logo" className="text-4xl" />
      <span className="mr-1 text-2xl font-barlow font-bold tracking-wider">LASSY SKY</span>
    </header>
  )
}

export default Logo