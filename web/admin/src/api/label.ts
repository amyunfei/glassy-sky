import { request } from '@/utils/http'

export interface Label {
  id: string
  name: string
  color: string
}

type CreateLabelParams = Omit<Label, 'id'>
export const createLabelApi = (data: CreateLabelParams) => request<Response<Label>>({
  url: '/label',
  method: 'POST',
  data
})

type UpdateLabelParams = Partial<CreateLabelParams>
export const updateLabelApi = (id: string, data: UpdateLabelParams) => request<Response<Label>>({
  url: `/label/${id}`,
  method: 'PUT',
  data
})

interface QueryLabelListParams {
  current: number
  size: number
}

export const queryLabelListApi = (data: QueryLabelListParams) => request<ResponseList<Label>>({
  url: '/label',
  method: 'GET',
  data
})

export const queryLabelDetailApi = (id: string) => request<Response<Label>>({
  url: `/label/${id}`,
  method: 'GET',
})

export const removeLabelApi = (id: string) => request<Response<null>>({
  url: `/label/${id}`,
  method: 'DELETE',
})