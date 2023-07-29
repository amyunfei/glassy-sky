import { useState, useImperativeHandle, forwardRef } from 'react'
import { Form, Input, message } from 'antd'
import { HexColorPicker } from 'react-colorful'
import { useTranslation } from 'react-i18next'
import { Label, createLabelApi, updateLabelApi } from '@/api/label'
import Dialog from '@/components/Dialog'

export type LabelEditorInstance = {
  open: (info?: Label) => void
}

const LabelEditor = forwardRef<LabelEditorInstance, {}>((_, ref) => {
  const [visible, setVisible] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)
  const [id, setId] = useState<string>('')
  const [title, setTitle] = useState<string>('')
  const [color, setColor] = useState<string>('#000000')
  const [form] = Form.useForm<{ name: string }>()


  useImperativeHandle(ref, () => ({
    open: (info?: Label) => {
      if (info) {
        setId(info.id)
        form.setFieldsValue({ name: info.name })
        setColor(info.color)
        setTitle(t('module.label.editLabel'))
      } else {
        setTitle(t('module.label.addLabel'))
      }
      setVisible(true)
    }
  }))
  const onConfirm = async () => {
    const { name } = await form.validateFields()
    let err
    setLoading(true)
    if (id) {
      [err] = await updateLabelApi(id, { name, color })
    } else {
      [err] = await createLabelApi({ name, color })
    }
    setLoading(false)
    if (err !== null) return
    if (id) {
      message.success(t('common-tips.updateSuccess'))
    } else {
      message.success(t('common-tips.createSuccess'))
    }
    onClose()
  }
  const onClose = () => {
    setVisible(false)
  }
  const afterClose = () => {
    form.resetFields()
    setColor('#000000')
  }

  const { t } = useTranslation()
  return (
    <Dialog
      open={visible}
      width={664}
      title={title}
      loading={loading}
      onConfirm={onConfirm}
      onClose={onClose}
      afterClose={afterClose}
    >
      <Form form={form} layout="vertical">
        <Form.Item name="name" label={t('module.label.name')} rules={[{ required: true, message: t('module.label.nameRequired') }]}>
          <Input placeholder={t('input-placeholder.input')} />
        </Form.Item>
        <Form.Item label={t('module.label.color')}>
          <div>
            <HexColorPicker color={color} onChange={setColor} style={{ width: '100%' }} />
            <div className="w-full h-6 flex items-center justify-end mt-3 pr-4 font-bold rounded" style={{ background: color }}>
              {color.toUpperCase()}
            </div>
          </div>
        </Form.Item>
      </Form>
    </Dialog>
  )
})

export default LabelEditor