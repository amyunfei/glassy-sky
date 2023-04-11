import React, { Suspense, CSSProperties } from 'react'
import { Outlet, useLocation } from 'react-router-dom'
import { useTransition, animated } from '@react-spring/web'
import { ConfigProvider, Layout as AntLayout } from 'antd'
import style from './layout.module.less'
import Logo from './Logo'
import Header from './Header'
import Sidebar from './Sidebar'
import Progress from '@/components/Progress'
const { Content: AntContent } = AntLayout

interface TransitionProps {
  children: React.ReactNode
}
const Transition: React.FC<TransitionProps> = ({ children }) => {
  const location = useLocation()
  const transitions = useTransition<string, CSSProperties>(location.pathname, {
    from: { opacity: 0, transform: 'translate3d(100%, 0, 0)' },
    enter: { opacity: 1, transform: 'translate3d(0%, 0, 0)' },
  })
  return transitions((style, item) => (
    <animated.div style={ style } key={ item } className="h-full w-full">
      { children }
    </animated.div>
  ))
}

const Layout: React.FC = () => {
  return (
    <AntLayout className="overflow-hidden flex-row h-screen">
      <div className="w-60 flex-shrink-0 bg-gradient-to-b from-gray-dark to-gray-light border-r border-solid border-color-base">
        <Logo className="bg-gray-dark" />
        <Sidebar />
      </div>
      <AntLayout className="main bg-gray-darker overflow-hidden">
        <Header />
        <Suspense fallback={ <Progress /> }>
          <AntContent data-simplebar className={ `${style['layout-content']} relative p-8` }>
            <div className="w-full h-full overflow-hidden">
              <ConfigProvider componentSize="large">
                <Transition>
                  <Outlet />
                </Transition>
              </ConfigProvider>
            </div>
          </AntContent>
        </Suspense>
      </AntLayout>
    </AntLayout>
  )
}

export default Layout