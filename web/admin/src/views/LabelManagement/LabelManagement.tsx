import React, { useState, useRef, useEffect } from 'react'
import type { ColumnsType } from 'antd/es/table'
import { Label, queryLabelListApi, queryLabelDetailApi, removeLabelApi } from '@/api/label'
import { deleteConfirm } from '@/utils/prompt'
import { useTranslation } from 'react-i18next'
import { Button, Space, message } from 'antd'
import TablePage from '@/components/TablePage'
import LabelEditor, { LabelEditorInstance } from './LabelEditor'


const UserManagement: React.FC = () => {
  const { t } = useTranslation()
  const [loading, setLoading] = useState<boolean>(false)
  const [dataSource, setDataSource] = useState<Label[]>([])
  const [total, setTotal] = useState<number>(0)
  const editorRef = useRef<LabelEditorInstance | null>(null)

  const fetchData = async () => {
    setLoading(true)
    const [err, res] = await queryLabelListApi({ size: 10, current: 1 })
    setLoading(false)
    if (err !== null) return
    setDataSource(res.data.list)
    setTotal(res.data.count)
  }
  useEffect(() => { fetchData() }, [])
  const handleEdit = (id: string) => {
    if (editorRef.current === null) return
    editorRef.current.open()
  }
  const handleDelete = (id: string) => {
    deleteConfirm(async () => {
      const [err, res] = await removeLabelApi(id)
      if (err === null) {
        message.success(t('common-tips.deleteSuccess'))
        fetchData()
      }
    })
  }
  const columns: ColumnsType<Label> = [
    { title: t('common-title.labelName'), dataIndex: 'name' },
    { title: t('common-title.color'), dataIndex: 'color', render: (_, { color }) => (
      <div className="w-10 h-6 rounded-md" style={{ backgroundColor: `#${color}` }} />
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
    <TablePage<Label>
      tableProps={{ columns, dataSource, loading, rowKey: 'id' }} total={total}
    >
      <LabelEditor ref={editorRef} />
    </TablePage>
  )
}

export default UserManagement