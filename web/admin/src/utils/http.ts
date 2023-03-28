import axios, { AxiosInstance, CreateAxiosDefaults, AxiosRequestConfig, AxiosError } from 'axios'
import { message } from 'antd'

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
        config.signal
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
        message.error(err.response.data.msg || err.message)
        return [err, undefined]
      })
  }
}

const httpRequest = new HttpRequest({
  baseURL: '/adminServer',
  timeout: 10000
})

export const request = httpRequest.request.bind(httpRequest)
