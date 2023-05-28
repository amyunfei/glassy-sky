import { useState, useEffect } from 'react'
import { AxiosError } from 'axios'

type Request<T, U> = (data: U) => Promise<[AxiosError<unknown, any>, undefined] | [null, ResponseList<T>]>
type OnPageChange = (current: number, pageSize: number) => void
interface PaginationProps {
  current: number
  pageSize: number
  total: number
}
interface TableProps<T> {
  dataSource: T[]
  loading: boolean
  pagination: PaginationProps
  onPageChange: OnPageChange
}
const useAntTable = <T, U>(request: Request<T, U>, defaultParams: U): TableProps<T> => {
  const [loading, setLoading] = useState(false)
  const [data, setData] = useState<T[]>([])
  const [pagination, setPagination] = useState<PaginationProps>({
    current: 1,
    pageSize: 10,
    total: 0
  })
  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await request(defaultParams)
    setLoading(false)
    if (err !== null) return
    setData(res.data.list)
    setPagination({ ...pagination, total: res.data.count })
  }
  const onPageChange: OnPageChange = (current, pageSize) => {
    console.log(current, pageSize)
    setPagination({ ...pagination, current, pageSize })
  }

  useEffect(() => {
    fetchData()
  }, [pagination.current, pagination.pageSize])

  return {
    dataSource: data,
    loading,
    pagination: pagination,
    onPageChange
  }
}

export default useAntTable