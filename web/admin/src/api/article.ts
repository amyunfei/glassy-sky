import { request } from '@/utils/http'

export interface Article {
  id: string
  title: string
  excerpt: string
  content: string
  category: string
  labels: string[]
  createdAt: string
  updatedAt: string
}

export const createArticle = (data: any) => request({
  url: '/article',
  method: 'POST',
  data
})