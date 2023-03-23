import { useState } from 'react'
import { Form, Input, Button } from 'antd'
import style from './Login.module.less'

const Cube: React.FC = () => {
  return (
    <div className={style.cube}>
      <div>GLASSY</div>
      <div>SKY</div>
      <div>ADMIN</div>
      <div></div>
    </div>
  )
}

const Login: React.FC = () => {
  const [loading, setLoading] = useState(false)
  return (
    <div className={style['page-login']}>
      <Cube />
      <Form className="w-[360px] py-8 px-6 bg-gray-dark absolute z-10 right-1/4 top-1/2 -translate-y-1/2 rounded-lg">
        <div className="text-xl mb-4 text-white font-bold">账号登录</div>
        <Form.Item className="mb-5">
          <Input placeholder="请输入账号" size="large"></Input>
        </Form.Item>
        <Form.Item className="mb-5">
          <Input.Password placeholder="请输入密码" size="large"></Input.Password>
        </Form.Item>
        <Button loading={loading} type="primary" size="large" block>登录</Button>
      </Form>
    </div>
  )
}

export default Login