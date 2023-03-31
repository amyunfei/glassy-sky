import 'i18next'
import en from './en-US.json'
import cn from './zh-CN.json'

declare module 'i18next' {
  interface CustomTypeOptions {
    defaultNS: 'en'
    resources: {
      en: typeof en
      cn: typeof cn
    }
  }
}