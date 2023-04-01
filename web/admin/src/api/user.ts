import { request } from '@/utils/http'
import { DtoLoginRequest, DtoCreateUserResponse } from './dto'

export const loginApi = (data: DtoLoginRequest) => request<DtoCreateUserResponse>({
  url: '/user/login',
  method: 'POST',
  data
})