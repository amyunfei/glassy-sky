import { request } from '@/utils/http'
import { DtoLoginRequest, DtoCreateUserResponse, DtoListResponseDtoCreateUserResponse, ResponseBodyString } from './dto'

export interface User {
  createdAt: string;
  email: string;
  id: string;
  nickname: string;
  updatedAt: string;
  username: string;
}


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