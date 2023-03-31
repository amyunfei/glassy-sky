import i18n, { Resource } from 'i18next'
import { initReactI18next } from 'react-i18next'
import cn from './zh-CN.json'
import en from './en-US.json'
import ja from './ja-JP.json'
export * from './Lang'
export { default } from './Lang'

const resources: Resource = {
  cn: {
    translation: cn,
  },
  en: {
    translation: en,
  },
  ja: {
    translation: ja,
  }
}

export const initI18n = (locale = 'cn') => {
  i18n.use(initReactI18next).init({
    resources,
    lng: locale,
    interpolation: {
      escapeValue: false
    }
  })
}

export const changeLanguage = (locale: string) => {
  return i18n.changeLanguage(locale)
}