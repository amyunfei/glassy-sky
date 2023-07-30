import { useState, forwardRef, useImperativeHandle } from 'react'
import { Divider, Form, Input, Row, Col, Upload, message } from 'antd'
import type { UploadChangeParam } from 'antd/es/upload'
import { UploadProps, UploadFile } from 'antd/es/upload/interface'
import { useTranslation } from 'react-i18next'
import Dialog from '@/components/Dialog'
import { User, CreateUserParams, createUserApi } from '@/api/user'
import { isEmpty, omit } from '@/utils/helper'

export type UserEditorInstance = {
  open: (info?: User) => void
}
const UserEditor = forwardRef<UserEditorInstance, {}>((_, ref) => {
  const [title, setTitle] = useState<string>()
  const [visible, setVisible] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)
  const [form] = Form.useForm<CreateUserParams>()
  const { t } = useTranslation()
  const [id, setId] = useState<string>('')

  useImperativeHandle(ref, () => ({
    open: (info?: User) => {
      if (!isEmpty(info)) {
        setId(info.id)
        form.setFieldsValue(omit(info, ['createdAt', 'updatedAt']))
        setTitle(t('module.user.editUser'))
      } else {
        setTitle(t('module.user.addUser'))
      }
      setVisible(true)
    }
  }))

  const onClose = () => {
    setVisible(false)
  }

  const afterClose = () => {
    form.resetFields()
  }

  const handleUploadChange: UploadProps['onChange'] = (info: UploadChangeParam<UploadFile>) => {
    const { file } = info
  }

  const onConfirm = async () => {
    try {
      const formData = await form.validateFields()
      setLoading(true)
      const [err] = await createUserApi(formData)
      setLoading(false)
      if (err !== null) return
      message.success(t('common-tips.createSuccess'))
      onClose()
    } catch (error) {
      console.warn(error)
    }
  }

  return (
    <Dialog
      width={1000}
      title={title}
      open={visible}
      loading={loading}
      onClose={onClose}
      onConfirm={onConfirm}
      afterClose={afterClose}
    >
      <Form form={form} layout="vertical">
        <Row>
          <Col span={5}>
            <Form.Item label="用户头像" className="flex justify-center">
              <Upload beforeUpload={() => false} onChange={handleUploadChange}>
                <div className="block w-48 pb-full bg-white rounded-full mt-4" />
              </Upload>
            </Form.Item>
          </Col>
          <Col span={1} className="text-center">
            <Divider type="vertical" className="h-full" />
          </Col>
          <Col span={18}>
            <Form.Item name="username" label="用户名" rules={[{ required: true, message: '请输入用户名' }]}>
              <Input />
            </Form.Item>
            <Form.Item name="password" label="密码" rules={[{ required: true, message: '请输入密码' }]}>
              <Input.Password />
            </Form.Item>
            <Form.Item name="nickname" label="昵称" rules={[{ required: true, message: '请输入昵称' }]}>
              <Input />
            </Form.Item>
            <Form.Item name="email" label="邮箱地址" rules={[{ required: true, message: '请输入邮箱地址' }]}>
              <Input />
            </Form.Item> 
          </Col>
        </Row>
      </Form>
    </Dialog>
  )
})

export default UserEditor