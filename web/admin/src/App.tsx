import { Router } from '@/router'
import { ConfigProvider } from 'antd'
import Empty from '@/components/Empty'

function App() :JSX.Element {
  return (
    <ConfigProvider renderEmpty={() => <Empty />}>
      <Router />
    </ConfigProvider>
  )
}

export default App