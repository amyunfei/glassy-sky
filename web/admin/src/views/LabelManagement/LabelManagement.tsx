import React, { useState, useEffect } from 'react'
import type { ColumnsType, TablePaginationConfig } from 'antd/es/table'
import { Label, queryLabelListApi } from '@/api/label'
import { useTranslation } from 'react-i18next'
import { Button, Space } from 'antd'
import TablePage from '@/components/TablePage'


const UserManagement: React.FC = () => {
  const [loading, setLoading] = useState<boolean>(false)
  const [dataSource, setDataSource] = useState<Label[]>([])
  const [pagination, setPagination] = useState<TablePaginationConfig>()

  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await queryLabelListApi({ size: 10, current: 1 })
    setLoading(false)
    if (err !== null) return
    setDataSource(res.data.list)
  }
  useEffect(() => { fetchData() }, [])
  const { t } = useTranslation()
  const columns: ColumnsType<Label> = [
    { title: t('common-title.labelName'), dataIndex: 'name' },
    { title: t('common-title.color'), dataIndex: 'color', render: (_, { color }) => (
      <div className="w-10 h-6 rounded-md" style={{ backgroundColor: `#${color}` }} />
    ) },
    { title: t('common-title.updatedAt'), dataIndex: 'updatedAt' },
    { title: t('common-title.createdAt'), dataIndex: 'createdAt' },
    { title: t('common-title.action'), dataIndex: 'action', render: (_, record) => (
      <Space size="large">
        <Button type="link" className="p-0">{t('common-action.edit')}</Button>
        <Button type="link" className="p-0">{t('common-action.delete')}</Button>
      </Space>
    ) }
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