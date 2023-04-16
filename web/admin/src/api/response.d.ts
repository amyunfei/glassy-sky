declare interface Response<T> {
  code: number
  message: string
  data: T
}

declare interface ResponseList<T> {
  code: number
  message: string
  data: {
    count: number
    list: T[]
  }
}