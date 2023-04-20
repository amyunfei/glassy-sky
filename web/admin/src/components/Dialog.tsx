import React from 'react'
import { Modal, Button, ModalFuncProps } from 'antd'
import { CheckOutlined, CloseOutlined, CloseSquareOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'


const bodyStyle = {
  padding: 0
}

interface DialogProps extends ModalFuncProps {
  children: React.ReactNode
  onClose?: (e: React.MouseEvent<HTMLElement, MouseEvent>) => void
  onConfirm?: (e: React.MouseEvent<HTMLElement, MouseEvent>) => void
  loading?: boolean
}
const Dialog: React.FC<DialogProps> = props => {
  const { t } = useTranslation()
  return (
    <Modal
      width="auto"
      centered
      closable={false}
      footer={null}
      bodyStyle={bodyStyle}
      {...props}
      title={null} 
    >
      {/* header */}
      <div className="px-8 h-20 flex justify-between items-center">
        <span className="text-2xl font-bold select-none">{ props.title }</span>
        <CloseSquareOutlined className="text-xl text-link" onClick={props.onClose} />
      </div>
      {/* content */}
      <div className="px-8 pb-10">
        { props.children }
      </div>
      {/* footer */}
      <div className="px-8 py-5 flex justify-end bg-black bg-opacity-25">
        <Button type="primary" loading={props.loading} icon={<CheckOutlined />} className="mr-3" onClick={props.onConfirm}>
          { t('common-action.submit') }
        </Button>
        <Button type="primary" ghost icon={<CloseOutlined />} onClick={props.onClose}>
          { t('common-action.cancel')}
        </Button>
      </div>
    </Modal>
  )
}

export default Dialog