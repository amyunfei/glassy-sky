import { Fragment, useEffect } from 'react'
import nprogress from 'nprogress'
import 'nprogress/nprogress.css'

const Progress = () => {
  useEffect(() => {
    nprogress.start()
    return () => {
      nprogress.done()
    }
  }, [])
  return <Fragment />
}

export default Progress