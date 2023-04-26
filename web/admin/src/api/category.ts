import { request } from '@/utils/http'

export interface Category {
  id: string
  name: string
  color: string
  parentId: string
  createdAt: string
  updatedAt: string
}

interface QueryCategoryListParams {
  current: number
  size: number
}

export const queryCategoryListApi = (data: QueryCategoryListParams) => request<ResponseList<Category>>({
  url: '/category',
  method: 'GET',
  data
})