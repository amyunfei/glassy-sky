import { makeAutoObservable } from 'mobx'

class AuthStore {
  token = ''

  constructor() {
    makeAutoObservable(this)
  }

  setToken(token: string) {
    this.token = token
  }
}

export default AuthStore