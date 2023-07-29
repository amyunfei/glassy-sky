import { request } from '@/utils/http'
import { DtoLoginRequest, ResponseBodyString } from './dto'

export interface User {
  id: string;
  email: string;
  username: string;
  nickname: string;
  createdAt: string;
  updatedAt: string;
}
export type CreateUserParams = Omit<User, 'id' | 'createdAt' | 'updatedAt'> & { password: string }

export const loginApi = (data: DtoLoginRequest) => request<ResponseBodyString>({
  url: '/user/login',
  method: 'POST',
  data
})

export const queryUserListApi = (data: any) => request<ResponseList<User>>({
  url: '/user',
  method: 'GET',
  data
})

export const removeUserApi = (id: string) => request<ResponseBodyString>({
  url: `/user/${id}`,
  method: 'DELETE'
})

export const createUserApi = (data: CreateUserParams) => request<Response<User>>({
  url: '/user',
  method: 'POST',
  data
})

export const queryUserApi = (id: string) => request<Response<User>>({
  url: `/user/${id}`,
  method: 'GET'
})