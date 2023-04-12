import { request } from '@/utils/http'

export interface Label {
  id: string
  name: string
  color: string
}
interface QueryLabelListParams {
  current: number
  size: number
}
interface QueryLabelListResponse {
  list: Label[]
  count: number
}
export const queryLabelListApi = (data: QueryLabelListParams) => request<Response<QueryLabelListResponse>>({
  url: '/label',
  method: 'GET',
  data
})

export const queryLabelDetailApi = (id: string) => request<Response<Label>>({
  url: `/label/${id}`,
  method: 'GET',
})