import React, { useState, useEffect } from 'react'
import type { ColumnsType, TablePaginationConfig } from 'antd/es/table'
import { queryUserListApi } from '@/api/user'
import { useTranslation } from 'react-i18next'
import { Button, Space } from 'antd'
import TablePage from '@/components/TablePage'


const UserManagement: React.FC = () => {
  const [loading, setLoading] = useState<boolean>(false)
  const [dataSource, setDataSource] = useState<any[]>([])
  const [pagination, setPagination] = useState<TablePaginationConfig>()

  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await queryUserListApi({ size: 10, current: 1 })
    setLoading(false)
    if (res && res.data && res.data.list) {
      setDataSource(res.data.list)
      setPagination({ total: res.data.count })
    }
  }
  useEffect(() => { fetchData() }, [])
  const { t } = useTranslation()
  const columns: ColumnsType<any> = [
    { title: t('common-title.username'), dataIndex: 'name' },
    { title: t('common-title.email'), dataIndex: 'email' },
    { title: t('common-title.phone'), dataIndex: 'phone' },
    { title: t('common-title.createdAt'), dataIndex: 'createdAt' },
    { title: t('common-title.action'), dataIndex: 'action', render: (_, record) => (
      <Space size="large">
        <Button type="link" className="p-0">{t('common-action.edit')}</Button>
        <Button type="link" className="p-0">{t('common-action.delete')}</Button>
      </Space>
    ), }
  ]

  return (
    <TablePage<any>
      tableProps={{ columns, dataSource, pagination, loading, rowKey: 'id' }}
    >
      {/*  */}
    </TablePage>
  )
}

export default UserManagement