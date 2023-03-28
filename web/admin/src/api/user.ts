import { request } from '@/utils/http'


export interface LoginApiRequest {
  username: string,
  password: string
}
export interface LoginApiResponse {
  code: 0,
  data: string,
  message: string
}
export const loginApi = (data: LoginApiRequest) => request<LoginApiResponse>({
  url: '/user/login',
  method: 'POST',
  data
})