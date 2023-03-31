import React from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Login from '@/views/Login'
import Layout from '@/layout'

const router = createBrowserRouter([
  {
    path: '/login',
    element: <Login />
  },
  {
    path: '/',
    element: <Layout />
  }
])

export const Router: React.FC = () => {
  return <RouterProvider router={router} />
}