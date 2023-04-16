import React from 'react'

interface CardProps extends React.HTMLAttributes<HTMLDivElement> {
  title: string
  children: React.ReactNode
}
const Card: React.FC<CardProps> = props => {
  return (
    <div className={`relative pt-15 rounded-md bg-gray-dark ${props.className || ''}`}>
      <div className="absolute top-0 inset-x-0 px-5 flex justify-between items-center h-15 border-b border-color-base">
        {/* decor */}
        <div className="absolute left-0 y-center w-1 h-6 bg-gradient-to-b from-blue-link to-blue rounded-r-sm" />
        {/* title */}
        <span className="text-base text-white font-semibold">{ props.title }</span>
      </div>
      { props.children }
    </div>
  )
}

export default Card