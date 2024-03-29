import React from 'react'
import {
  FundProjectionScreenOutlined, TagOutlined, AppstoreOutlined, BookOutlined, UnorderedListOutlined, CompassOutlined,
  UserOutlined, FileTextOutlined
} from '@ant-design/icons'
import Login from '@/views/Login'
import Layout from '@/layout'
import NotFound from '@/views/NotFound'

export const menuRoutes: AppRoute.Route[] = [
  {
    path: 'dashboard',
    name: 'Dashboard',
    component: React.lazy(() => import('@/views/Dashboard/Dashboard')),
    meta: { title: 'Dashboard', icon: FundProjectionScreenOutlined }
  },
  {
    path: 'user-management',
    name: 'UserManagement',
    component: React.lazy(() => import('@/views/UserManagement')),
    meta: { title: 'User-Management', icon: UserOutlined, breadcrumb: true }
  },
  {
    path: 'category-management',
    name: 'CategoryManagement',
    component: React.lazy(() => import('@/views/CategoryManagement')),
    meta: { title: 'Category-Management', icon: AppstoreOutlined, breadcrumb: true }
  },
  {
    path: 'label-management',
    name: 'LabelManagement',
    component: React.lazy(() => import('@/views/LabelManagement')),
    meta: { title: 'Label-Management', icon: TagOutlined, breadcrumb: true }
  },
  {
    path: 'article-management',
    name: 'ArticleManagement',
    component: React.lazy(() => import('@/views/ArticleManagement')),
    meta: { title: 'Article-Management', icon: FileTextOutlined, breadcrumb: true }
  }
  // {
  //   path: 'exercise',
  //   name: 'Exercise',
  //   redirect: '/exercise/category',
  //   meta: { title: 'Exercise', icon: CoffeeOutlined },
  //   children: [
  //     {
  //       path: 'categories',
  //       name: 'Categories',
  //       component: React.lazy(() => import('@/views/Exercise/Category')),
  //       meta: { title: 'Categories', breadcrumb: true }
  //     }
  //   ]
  // },
  // {
  //   path: 'list',
  //   name: 'List',
  //   redirect: '/list/table-list',
  //   meta: { title: 'List', icon: UnorderedListOutlined },
  //   children: [
  //     {
  //       path: 'table-list',
  //       name: 'TableList',
  //       component: React.lazy(() => import('@/views/List/table-list')),
  //       meta: { title: 'List-Table', breadcrumb: true }
  //     },
  //     {
  //       path: 'card-list',
  //       name: 'CardList',
  //       component: React.lazy(() => import('@/views/List/card-list')),
  //       meta: { title: 'List-Card', breadcrumb: true }
  //     }
  //   ]
  // },
  // {
  //   path: 'map',
  //   name: 'Map',
  //   component: React.lazy(() => import('@/views/Map')),
  //   meta: { title: 'Map', icon: CompassOutlined }
  // },
  // {
  //   path: 'categories',
  //   name: 'Categories',
  //   component: React.lazy(() => import('../views/Categories/index')),
  //   meta: { title: '分类', icon: AppstoreOutlined, breadcrumb: true }
  // },
  // {
  //   path: 'notes',
  //   name: 'Notes',
  //   component: React.lazy(() => import('../views/Notes/Notes')),
  //   meta: { title: 'Notes', icon: BookOutlined, breadcrumb: true }
  // }
]

export const staticRoutes: AppRoute.Route[] = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { title: 'Login', hidden: true }
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFound,
    meta: { title: 'NotFound', hidden: true }
  }
]

export const dynamicRoutes: AppRoute.Route[] = [
  {
    path: '/',
    name: 'Layout',
    component: Layout,
    meta: { title: 'Layout', hidden: true },
    redirect: '/dashboard',
    children: menuRoutes
  }
]