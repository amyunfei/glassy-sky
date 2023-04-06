import { useState } from 'react'
import { Form, Input, Button, message } from 'antd'
import style from './Login.module.less'
import { useStore } from '@/store'
import { loginApi } from '@/api/user'
import { DtoLoginRequest } from '@/api/dto'

const Cube: React.FC = () => {
  return (
    <div className={style.cube}>
      <div>GLASSY</div>
      <div>SKY</div>
      <div>ADMIN</div>
      <div />
    </div>
  )
}

const Login: React.FC = () => {
  const [loading, setLoading] = useState(false)
  const { auth } = useStore()

  const onFinish = async (values: DtoLoginRequest) => {
    setLoading(true)
    const [err, res] = await loginApi(values)
    setLoading(false)
    if (err !== null) return
    auth.setToken(res.data || '')
    message.success('登录成功')
  }

  return (
    <div className={style['page-login']}>
      <Cube />
      <Form<DtoLoginRequest>
        onFinish={onFinish}
        className="w-[360px] py-8 px-6 bg-gray-dark absolute z-10 right-1/4 top-1/2 -translate-y-1/2 rounded-lg"
      >
        <div className="text-xl mb-4 text-white font-bold">账号登录</div>
        <Form.Item
          name="username"
          rules={[{ required: true, message: '账号不能为空' }]}
          className="mb-6"
        >
          <Input placeholder="请输入账号" size="large" />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: '密码不能为空' }]}
          className="mb-6"
        >
          <Input.Password placeholder="请输入密码" size="large" />
        </Form.Item>
        <Button loading={loading} type="primary" htmlType="submit" size="large" block>登录</Button>
      </Form>
    </div>
  )
}

export default Login