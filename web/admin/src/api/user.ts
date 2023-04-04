import { request } from '@/utils/http'
import { DtoLoginRequest, DtoCreateUserResponse, ResponseBodyDtoListResponseDtoCreateUserResponse } from './dto'

export const loginApi = (data: DtoLoginRequest) => request<DtoCreateUserResponse>({
  url: '/user/login',
  method: 'POST',
  data
})

export const queryUserListApi = (data: any) => request<ResponseBodyDtoListResponseDtoCreateUserResponse>({
  url: '/user',
  method: 'GET',
  data
})