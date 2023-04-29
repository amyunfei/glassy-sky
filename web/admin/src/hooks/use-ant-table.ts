import { useState, useEffect } from 'react'
import { TableProps } from 'antd'
import { AxiosError } from 'axios'
import type { TablePaginationConfig } from 'antd/es/table'

type Request<T, U> = (data: U) => Promise<[AxiosError<unknown, any>, undefined] | [null, ResponseList<T>]>
const useAntTable = <T, U>(request: Request<T, U>, defaultParams: U) => {
  const [loading, setLoading] = useState(false)
  const [data, setData] = useState<T[]>([])
  const [pagination, setPagination] = useState<TablePaginationConfig>({
    current: 1,
    pageSize: 10
  })
  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await request(defaultParams)
    setLoading(false)
    if (err !== null) return
    setData(res.data.list)
  }

  useEffect(() => {
    fetchData()
  }, [])

  const tableProps: TableProps<T> = {
    dataSource: data,
    loading: loading
  }
  return { tableProps }
}

export default useAntTable