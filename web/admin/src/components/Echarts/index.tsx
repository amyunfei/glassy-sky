import { useRef, useEffect } from 'react'
import * as echarts from 'echarts'
import { useEventListener } from '@/hooks'

interface EchartsProps extends React.HTMLAttributes<HTMLDivElement> {
  option: echarts.EChartsOption
}
const Echarts: React.FC<EchartsProps> = props => {
  const renderer = useRef<HTMLDivElement | null>(null)
  let echartsInstance: echarts.ECharts | null = null

  useEffect(() => {
    if (renderer.current) {
      echartsInstance = echarts.init(renderer.current, props.option)
    }
    return () => {
      echartsInstance?.dispose()
    }
  }, [renderer.current, props.option])

  useEventListener('resize', () => { echartsInstance?.resize() })
  console.log('echarts render')
  return (
    <div ref={renderer} className={props.className} />
  )
}

export default Echarts