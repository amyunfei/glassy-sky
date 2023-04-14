import { useState, useImperativeHandle, forwardRef } from 'react'
import { Form, Input } from 'antd'
import { RgbColorPicker } from 'react-colorful'
import { useTranslation } from 'react-i18next'
import Dialog from '@/components/Dialog'

export type LabelEditorInstance = {
  open: () => void
}

const LabelEditor = forwardRef<LabelEditorInstance, {}>((_, ref) => {
  const [visible, setVisible] = useState<boolean>(false)
  const [form] = Form.useForm()
  useImperativeHandle(ref, () => ({
    open: () => {
      setVisible(true)
    }
  }))
  const onClose = () => {
    setVisible(false)
    form.resetFields()
  }

  const { t } = useTranslation()
  return (
    <Dialog open={visible} width={664} title={t('module.label.addLabel')} onClose={onClose}>
      <Form form={form} layout="vertical">
        <Form.Item label={t('module.label.name')} name="name">
          <Input placeholder={t('input-placeholder.input')} />
        </Form.Item>
        <Form.Item label={t('module.label.color')} name="color">
          <RgbColorPicker className="w-full" />
          <div className="w-full h-6 bg-indigo-50 rounded mt-3" />
        </Form.Item>
      </Form>
    </Dialog>
  )
})

export default LabelEditor