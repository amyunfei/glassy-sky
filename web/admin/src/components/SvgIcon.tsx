
interface propsType extends React.HTMLAttributes<SVGSVGElement> {
  name: string
  prefix?: string
  color?: string
}

function SvgIcon ({ name, prefix = 'icon', color = '#333', ...props }: propsType): JSX.Element {
  const symbolId = `#${prefix}-${name}`
  return (
    <svg {...props} aria-hidden="true" className={'svg-icon ' + (props.className || '')}>
      <use href={symbolId} fill={color} />
    </svg>
  )
}

export default SvgIcon