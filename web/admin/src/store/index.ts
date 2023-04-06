import React from 'react'
import AuthStore from './auth'

class RootStore {
  auth: AuthStore
  constructor() {
    this.auth = new AuthStore()
  }
}

const storeContext = React.createContext(new RootStore())
const useStore = () => React.useContext(storeContext)
export { storeContext, useStore }