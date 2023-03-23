import { BrowserRouter } from 'react-router-dom'
import Login from '@/views/Login/Login'
import { ConfigProvider } from 'antd'

function App() :JSX.Element {
  return (
    <ConfigProvider>
      <Login></Login>
      {/* <BrowserRouter></BrowserRouter> */}
    </ConfigProvider>
  )
}

export default App