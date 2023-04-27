import React, { useEffect, useState } from 'react'
import { Button, Space } from 'antd'
import type { ColumnsType } from 'antd/es/table'
import { useTranslation } from 'react-i18next'
import { Category, queryCategoryListApi } from '@/api/category'
import TablePage from '@/components/TablePage'



const CategoryManagement: React.FC = () => {
  const { t } = useTranslation()
  const [loading, setLoading] = useState<boolean>(false)
  const [listQuery, setListQuery] = useState<{ size: number, current: number }>({ size: 10, current: 1 })
  const [dataSource, setDataSource] = useState<Category[]>()
  const [total, setTotal] = useState<number>(0)

  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await queryCategoryListApi(listQuery)
    setLoading(false)
    if (err !== null) return
    setDataSource(res.data.list)
    setTotal(res.data.count)
  }
  const handlePageChange = (current: number, size: number) => {
    setListQuery({ size, current })
  }
  const handleEdit = async (id?: string) => {
    // 
  }
  const handleDelete = (id: string) => {
    // 
  }

  useEffect(() => { fetchData() }, [listQuery])

  const columns: ColumnsType<Category> = [
    { title: t('module.category.name'), dataIndex: 'name' },
    { title: t('common-title.color'), dataIndex: 'color', render: (_, { color }) => (
      <div className="w-10 h-6 rounded-md" style={{ backgroundColor: color }} />
    ) },
    { title: t('common-title.updatedAt'), dataIndex: 'updatedAt' },
    { title: t('common-title.createdAt'), dataIndex: 'createdAt' },
    { title: t('common-title.action'), dataIndex: 'action', render: (_, record) => (
      <Space size="large">
        <Button type="link" className="p-0" onClick={() => handleEdit(record.id)}>{t('common-action.edit')}</Button>
        <Button type="link" className="p-0" onClick={() => handleDelete(record.id)}>{t('common-action.delete')}</Button>
      </Space>
    ) }
  ]

  return (
    <TablePage<Category>
      tableProps={{ columns, dataSource, loading, rowKey: 'id' }}   
      current={listQuery.current}
      pageSize={listQuery.size}
      total={total}
      pageChange={handlePageChange}
    >
      {/*  */}
    </TablePage>
  )
}

export default CategoryManagement