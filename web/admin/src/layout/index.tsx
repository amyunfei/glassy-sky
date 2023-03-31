import React, { Suspense } from 'react'
import { Outlet } from 'react-router-dom'
import { ConfigProvider, Layout as AntLayout } from 'antd'
import Logo from './Logo'
import Header from './Header'
import Progress from '@/components/Progress'

const { Content } = AntLayout
const Layout: React.FC = () => {
  return (
    <AntLayout className="overflow-hidden flex-row h-screen">
      <div className="w-60 flex-shrink-0 bg-gradient-to-b from-gray-dark to-gray-light border-r border-solid border-color-base">
        <Logo className="bg-gray-dark" />
        {/* <Logo className='bg-gray-dark' />
          <Sidebar /> */}
      </div>
      <AntLayout className="main bg-gray-darker">
        <Header />
        <Suspense fallback={ <Progress /> }>
          <Content data-simplebar className="layout-content relative p-8">
            <div className="w-full h-full">
              <ConfigProvider componentSize="large">
                {/* <PageTitle /> */}
                { new Array(100).fill(null).map(item => {
                  return <div className="h-10 bg-gray-dark mb-2" />
                }) }
                <Outlet />
              </ConfigProvider>
            </div>
          </Content>
        </Suspense>
      </AntLayout>
    </AntLayout>
  )
}

export default Layout