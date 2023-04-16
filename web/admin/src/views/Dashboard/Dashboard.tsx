import React from 'react'
import { Row, Col } from 'antd'
import Card from '@/components/Card'
import Echarts from '@/components/Echarts'
import GithubContributions from './GithubContributions'


const Dashboard: React.FC = () => {
  return (
    <Row gutter={[32, 32]}>
      <Col xl={24}>
        <Card title="Github Contributions" className="h-96">
          <GithubContributions />
        </Card>
      </Col>
    </Row>
  )
}

export default Dashboard