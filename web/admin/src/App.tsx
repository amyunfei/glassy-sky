import { BrowserRouter } from 'react-router-dom'
import { ConfigProvider } from 'antd'

function App() :JSX.Element {
  return (
    <ConfigProvider>
      <BrowserRouter></BrowserRouter>
    </ConfigProvider>
  )
}

export default App