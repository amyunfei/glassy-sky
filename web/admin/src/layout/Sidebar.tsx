import React, { useMemo } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { Menu, MenuProps } from 'antd'
import { menuRoutes } from '@/router/routes'
type MenuItem = Required<MenuProps>['items'][number]

function generateMenu (routes: AppRoute.Route[], prefix: string = '') {
  const arr: MenuItem[] = []
  const len = routes.length
  for (let i = 0; i < len; i++) {
    const route = routes[i]
    const path = prefix + '/' + route.path
    const icon = route.meta.icon ? <route.meta.icon /> : ''
    if (route.meta.hidden !== true) { // 路由在菜单显示
      const children = route.children
      if (children !== undefined && children.length > 0) { // 存在子路由
        const childrenAllHidden = children.every(item => item.meta.hidden === true)
        if (childrenAllHidden) {
          arr.push(
            { key: path, icon: icon, label: <Link to={path}>{route.meta.title}</Link> }
          )
        } else {
          arr.push(
            { key: path, icon: icon, label: route.meta.title, children: generateMenu(children, path) }
          )
        }
      } else {
        arr.push(
          { key: path, icon: icon, label: <Link to={path}>{route.meta.title}</Link> }
        )
      }
    }
  }
  return arr
}

function useMenu () {
  const { pathname } = useLocation()
  const result = useMemo(() => {
    return [pathname]
  }, [pathname])
  return result
}


const Sidebar: React.FC = () => {
  const [pathname] = useMenu()
  return (
    <Menu
      items={generateMenu(menuRoutes)}
      mode="inline"
      selectedKeys={[pathname]}
      className="bg-transparent"
    />
  )
}

export default Sidebar