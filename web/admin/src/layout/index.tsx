import React, { Suspense } from 'react'
import { Outlet } from 'react-router-dom'
import { ConfigProvider, Layout as AntLayout } from 'antd'
import Logo from './Logo'
import Header from './Header'
import Sidebar from './Sidebar'
import Progress from '@/components/Progress'
const { Content: AntContent } = AntLayout

const Content: React.FC = () => {
  return (
    <AntContent data-simplebar className="layout-content relative p-8">
      <div className="w-full h-full overflow-hidden">
        <ConfigProvider componentSize="large">
          {/* <PageTitle /> */}
          <Outlet />
        </ConfigProvider>
      </div>
    </AntContent>
  )
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
          <Content />
        </Suspense>
      </AntLayout>
    </AntLayout>
  )
}

export default Layout