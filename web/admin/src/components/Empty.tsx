import React from 'react'
import SvgIcon from './SvgIcon'
import { useTranslation } from 'react-i18next'

const Empty: React.FC = () => {
  const { t } = useTranslation()
  return (
    <div>
      <SvgIcon name="data_empty" className="text-9xl" />
      <div className="font-bold">{t('common-tips.noData')}</div>
    </div>
  )
}

export default Empty