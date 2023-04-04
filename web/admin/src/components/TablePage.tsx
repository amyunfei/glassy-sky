import React, { useState, useEffect, useRef } from 'react'
import { Table, TableProps, Button } from 'antd'
import { PlusOutlined } from '@ant-design/icons'

interface PropsType<T> {
  handleAdd?: () => void,
  tableProps: TableProps<T>
  children?: React.ReactNode
}

function TablePage<T extends object> (props: PropsType<T>) {
  const tableRef = useRef<HTMLDivElement>(null)
  const [tableTop, setTableTop] = useState<number>(0)

  useEffect(() => {
    if (tableRef.current !== null) {
      setTableTop(tableRef.current.getBoundingClientRect().top)
    }
  }, [])
  return (
    <div className="h-full flex flex-col">
      <div className="flex items-center flex-shrink-0 mb-5">
        <Button type="primary" icon={<PlusOutlined />} className="ml-auto" onClick={props.handleAdd} />
      </div>
      <Table
        {...props.tableProps}
        ref={tableRef}
        scroll={{ y: `calc(100vh - ${tableTop}px - 64px - 55px - 32px)` }}
        className="flex-grow h-auto" />
      { props.children }
    </div>
  )
}

export default TablePage