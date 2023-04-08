import React, { useState, useEffect, useRef } from 'react'
import { Table, TableProps, Pagination, Button } from 'antd'
import { PlusOutlined } from '@ant-design/icons'

interface PropsType<T> {
  handleAdd?: () => void,
  tableProps: TableProps<T>
  children?: React.ReactNode
}

function TablePage<T extends object> (props: PropsType<T>) {
  const tableRef = useRef<HTMLDivElement>(null)
  const [contentHeight, setContentHeight] = useState<number>(0)

  useEffect(() => {
    if (tableRef.current !== null) {
      const thead = tableRef.current.querySelector('.ant-table-thead')
      const tablePlaceholder = tableRef.current.querySelector<HTMLElement>('.ant-table-placeholder')
      if (thead !== null && tablePlaceholder !== null) {
        const contentHeight = tableRef.current.getBoundingClientRect().height - thead.getBoundingClientRect().height
        tablePlaceholder.style.height = `${contentHeight}px`
        setContentHeight(contentHeight)
      }
    }
  }, [])
  return (
    <div className="h-full flex flex-col">
      <div className="flex items-center flex-shrink-0 mb-5">
        <Button type="primary" icon={<PlusOutlined />} className="ml-auto" onClick={props.handleAdd} />
      </div>
      <Table
        {...props.tableProps}
        pagination={false}
        ref={tableRef}
        scroll={{ y: `${contentHeight}px` }}
        className="flex-grow h-auto"
      />
      <Pagination className="text-right pt-4 flex-shrink-0" />
      { props.children }
    </div>
  )
}

export default TablePage