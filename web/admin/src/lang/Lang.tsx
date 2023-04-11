import React, { useState } from 'react'
import { changeLanguage } from '@/lang'
import { Dropdown, MenuProps } from 'antd'
import type { ItemType } from 'antd/es/menu/hooks/useItems'
import SvgIcon from '@/components/SvgIcon'

const languages = [
  { key: 'cn', label: '简体中文', icon: 'flag_cn' },
  { key: 'en', label: 'Englist', icon: 'flag_en' },
  { key: 'ja', label: '日本語', icon: 'flag_ja' }
]
const LangItems: ItemType[] = languages.map(language => {
  return {
    key: language.key,
    label: language.label,
    icon: <SvgIcon name={language.icon} className="rounded mr-2" />
  }
})

const Lang: React.FC = () => {
  const [langIcon, setLangIcon] = useState<string>('flag_en')
  const onClick: MenuProps['onClick'] = ({ key }) => {
    changeLanguage(key)
    setLangIcon('flag_' + key)
  }
  return (
    <Dropdown menu={{ items: LangItems, onClick }} placement="bottom">
      <SvgIcon name={langIcon} className="text-2xl rounded cursor-pointer" />
    </Dropdown>
  )
}

export default Lang