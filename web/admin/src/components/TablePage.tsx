import React, { useState, useEffect, useRef, useCallback } from 'react'
import { Table, TableProps, Pagination, Button } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import { useEventListener } from '@/hooks'

interface PropsType<T> {
  handleAdd?: () => void,
  tableProps: TableProps<T>
  children?: React.ReactNode
  total: number
}

function TablePage<T extends object> (props: PropsType<T>) {
  const tableRef = useRef<HTMLDivElement>(null)
  const [contentHeight, setContentHeight] = useState<number>(0)

  const computedTableHeight = useCallback(() => {
    if (tableRef.current !== null) {
      const thead = tableRef.current.querySelector('.ant-table-thead')
      const tablePlaceholder = tableRef.current.querySelector<HTMLElement>('.ant-table-placeholder')
      setTimeout(() => {
        if (thead !== null && tablePlaceholder !== null && tableRef.current !== null) {
          const contentHeight = tableRef.current.getBoundingClientRect().height - thead.getBoundingClientRect().height
          tablePlaceholder.style.height = `${contentHeight}px`
          setContentHeight(contentHeight)
        }
      }, 30)
    }
  }, [])

  useEffect(() => { computedTableHeight() }, [])
  useEventListener('resize', computedTableHeight)
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
        className="flex-grow h-auto overflow-hidden"
      />
      <Pagination total={props.total} className="text-right pt-4 flex-shrink-0" />
      { props.children }
    </div>
  )
}

export default TablePage