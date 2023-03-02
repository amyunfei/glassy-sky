import React from 'react'
import ReactDOM from 'react-dom/client'
import './styles/normalize.css'
import 'virtual:svg-icons-register'

import 'simplebar'
import 'simplebar/dist/simplebar.css'

import './styles/antd-theme.less'
import './index.css'

import App from './App'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
