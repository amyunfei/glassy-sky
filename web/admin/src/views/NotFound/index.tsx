import React from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Result } from 'antd'
import SvgIcon from '@/components/SvgIcon'

const NotFount: React.FC = () => {
  const navigate = useNavigate()
  return (
    <Result
      icon={<SvgIcon name="404" className="w-96 h-72" />}
      subTitle="Sorry, the page you visited does not exist."
      extra={<Button type="primary" onClick={() => navigate('/')}>Back Home</Button>}
      className="h-screen w-screen flex flex-col justify-center"
    />
  )
}

export default NotFount