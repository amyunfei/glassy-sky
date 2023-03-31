import { Router } from '@/routes'
import { ConfigProvider } from 'antd'
import Empty from '@/components/Empty'

function App() :JSX.Element {
  return (
    <ConfigProvider renderEmpty={() => <Empty />}>
      <Router></Router>
    </ConfigProvider>
  )
}

export default App