declare namespace AppRoute {
  type Route = {
    path: string,
    name: string,
    component?: any,
    meta: RouteMeta,
    redirect?: string,
    children?: Route[]
  }

  type RouteMeta = {
    title?: string,
    icon?: React.ForwardRefExoticComponent<any>,
    hidden?: boolean,
    breadcrumb?: boolean
  }
}