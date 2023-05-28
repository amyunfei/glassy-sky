import React, { useEffect, useState } from 'react'
import { Button, Space } from 'antd'
import type { ColumnsType } from 'antd/es/table'
import { useTranslation } from 'react-i18next'
import { Category, queryCategoryListApi } from '@/api/category'
import useAntTable from '@/hooks/use-ant-table'
import TablePage from '@/components/TablePage'



const CategoryManagement: React.FC = () => {
  const { t } = useTranslation()
  const { dataSource, loading, pagination, onPageChange } = useAntTable(queryCategoryListApi, { size: 10, current: 1 })
  const handleEdit = async (id?: string) => {
    // 
  }
  const handleDelete = (id: string) => {
    // 
  }

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
      tableProps={{ columns, rowKey: 'id', loading, dataSource }}   
      current={pagination.current}
      pageSize={pagination.pageSize}
      total={pagination.total}
      pageChange={onPageChange}
    >
      {/*  */}
    </TablePage>
  )
}

export default CategoryManagement