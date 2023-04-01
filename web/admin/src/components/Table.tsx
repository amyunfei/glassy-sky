import React from 'react'
import { Table as AntTable, TableProps, Button } from 'antd'
import { PlusOutlined } from '@ant-design/icons'

interface PropsType<T> {
  handleAdd?: () => void,
  tableProps: TableProps<T>
  children?: React.ReactNode
}

function Table<T extends object> (props: PropsType<T>) {
  return (
    <div>
      <div className="flex items-center mb-5">
        <Button type="primary" icon={<PlusOutlined />} className="ml-auto" onClick={props.handleAdd} />
      </div>
      <AntTable {...props.tableProps} />
      { props.children }
    </div>
  )
}

export default Table