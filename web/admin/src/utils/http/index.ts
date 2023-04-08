import axios, { AxiosInstance, CreateAxiosDefaults, AxiosRequestConfig, AxiosError } from 'axios'
import { message } from 'antd'
import i18n from 'i18next'
import { HttpStatusCode } from './http-status-code'

interface RequestConfig {
  url: string,
  method: 'GET' | 'POST' | 'PUT' | 'DELETE',
  data?: any,
}
export default class HttpRequest {
  instance: AxiosInstance
  constructor(config?: CreateAxiosDefaults) {
    this.instance = axios.create(config)

    this.instance.interceptors.request.use(
      config => {
        const token = localStorage.getItem('token')
        config.headers.Authorization = `Bearer ${token}`
        return config
      },
      error => {
        return Promise.reject(error)
      }
    )

    this.instance.interceptors.response.use(
      response => {
        return response
      },
      err => {
        const errResponse = (err as AxiosError).response
        if (errResponse && errResponse.status === HttpStatusCode.UNAUTHORIZED) {
          message.error(i18n.t('network.unAuthorized'))
          localStorage.removeItem('token')
        }
        return Promise.reject(err)
      }
    )
  }

  request<T>(option: RequestConfig): Promise<[null, T] | [AxiosError, undefined]> {
    const config: AxiosRequestConfig = { ...option }
    if (option.method === 'GET') {
      config.params = option.data
      delete config.data
    }
    return this.instance<T>(config)
      .then<[null, T]>(res => [null, res.data])
      .catch<[AxiosError, undefined]>(err => {
        return [err, undefined]
      })
  }
}

const httpRequest = new HttpRequest({
  baseURL: '/adminServer',
  timeout: 10000
})

export const request = httpRequest.request.bind(httpRequest)
