import { useRef } from 'react'
import type { ColumnsType } from 'antd/es/table'
import useAntTable from '@/hooks/use-ant-table'
import { User, queryUserListApi, queryUserApi, removeUserApi } from '@/api/user'
import { useTranslation } from 'react-i18next'
import { Button, Space, message } from 'antd'
import UserEditor, { UserEditorInstance } from './UserEditor'
import TablePage from '@/components/TablePage'
import { deleteConfirm } from '@/utils/prompt'


const UserManagement: React.FC = () => {
  const editorRef = useRef<UserEditorInstance | null>(null)
  const { t } = useTranslation()
  const {
    fetchData, dataSource, loading, pagination, onPageChange
  } = useAntTable(queryUserListApi, { current: 1, size: 10 })

  const handleEdit = async (id?: string) => {
    if (editorRef.current === null) return
    let info: User | undefined
    if (id !== undefined) {
      const [err, res] = await queryUserApi(id)
      if (err !== null) return
      info = res.data
    }
    editorRef.current.open(info)
  }

  const handleDelete = (id: string) => {
    deleteConfirm(async () => {
      const [err] = await removeUserApi(id)
      if (err === null) {
        message.success(t('common-tips.deleteSuccess'))
        fetchData()
      }
    })
  }


  const columns: ColumnsType<User> = [
    { title: t('common-title.username'), dataIndex: 'username' },
    { title: t('common-title.nickname'), dataIndex: 'nickname' },
    { title: t('common-title.email'), dataIndex: 'email' },
    { title: t('common-title.updatedAt'), dataIndex: 'updatedAt' },
    { title: t('common-title.createdAt'), dataIndex: 'createdAt' },
    { title: t('common-title.createdAt'), dataIndex: 'createdAt' },
    { title: t('common-title.action'), dataIndex: 'action', render: (_, record) => (
      <Space size="large">
        <Button type="link" className="p-0" onClick={() => handleEdit(record.id)}>{t('common-action.edit')}</Button>
        <Button type="link" className="p-0" onClick={() => handleDelete(record.id)}>{t('common-action.delete')}</Button>
      </Space>
    ) }
  ]
  return (
    <TablePage<User>
      tableProps={{ columns, rowKey: 'id', loading, dataSource }}
      total={pagination.total}
      pageSize={pagination.pageSize}
      current={pagination.current}
      pageChange={onPageChange}
      handleAdd={() => handleEdit()}
    >
      <UserEditor ref={editorRef} />
    </TablePage>
  )
}

export default UserManagement