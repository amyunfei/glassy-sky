import { Modal } from 'antd'
import i18n from 'i18next'

type DeleteFn = () => Promise<void>
export const deleteConfirm = (onOk: DeleteFn, title?: string) => {
  Modal.confirm({
    title: title || i18n.t('common-tips.deleteConfirm'),
    centered: true,
    okText: i18n.t('common-action.confirm'),
    okType: 'danger',
    cancelText: i18n.t('common-action.cancel'),
    onOk
  })
}