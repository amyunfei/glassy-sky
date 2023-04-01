import React from 'react'
import { BrowserRouter, Routes, Route, Navigate, Outlet } from 'react-router-dom'
import { routes } from './routes' 


function generateRoutes (routes: AppRoute.Route[]): (JSX.Element | undefined)[] {
  return routes.map(route => {
    if (route.children !== undefined && route.children.length > 0) {
      let routeElement, redirectRoute: (undefined | React.ReactElement)
      
      if (route.component) {
        routeElement = <route.component />
        if (route.redirect !== undefined) {
          redirectRoute = <Route key={route.path + '/*'} index element={<Navigate to={route.redirect}/>} />
        }
      } else if (route.redirect !== undefined) {
        redirectRoute = <Route key={route.path + '/*'} index element={<Navigate to={route.redirect}/>} />
        return (
          <Route key={ route.path } path={ route.path } element={ <Outlet /> }>
            { redirectRoute }
            { generateRoutes(route.children) }
          </Route>
        )
      }
      return (
        <Route key={ route.path } path={ route.path } element={ routeElement }>
          { redirectRoute }
          { generateRoutes(route.children) }
        </Route>
      )
    } else {
      return <Route key={ route.path } path={ route.path } element={ <route.component /> } />
    }
  })
}

export const Router: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        { generateRoutes(routes)}
      </Routes>
    </BrowserRouter>
  )
}