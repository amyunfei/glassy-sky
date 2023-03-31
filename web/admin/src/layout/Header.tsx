import { Input, Avatar, Badge, Divider } from 'antd'
import { SearchOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import Lang from '@/lang'
const Header: React.FC = () => {
  const { t } = useTranslation()

  return (
    <header
      className="h-20 px-8 flex justify-between items-center flex-shrink-0 bg-gray-dark"
      style={{ boxShadow: '0px 15px 10px -15px #000', zIndex: 999 }}
    >
      <span className="text-base text-gray-lightest font-barlow font-bold tracking-wider">{ t('navbar.welcome') }</span>
      <div className="flex items-center">
        <Input
          size="large"
          placeholder={t('input-placeholder.search')}
          prefix={<SearchOutlined className="text-gray-lightest" />}
          className="w-70 bg-gray-darker"
        />
        <Divider type="vertical" className="h-5 mx-6 bg-gray-lightest" />
        {/* <Badge dot={true} color="green" className=" mr-5">
          <MailOutlined className="text-gray-lightest font-black text-lg leading-4" />
        </Badge>
        <Badge dot={true} color="red">
          <BellOutlined className="text-gray-lightest font-black text-xl leading-4" />
        </Badge> */}
        <Badge dot={true} color="green">
          <Avatar shape="square" />
        </Badge>
        <Divider type="vertical" className="h-5 mx-6 bg-gray-lightest" />
        <Lang />
      </div>
    </header>
  )
}

export default Header