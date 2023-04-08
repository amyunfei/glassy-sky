import { makeAutoObservable } from 'mobx'

export class AuthStore {
  token = localStorage.getItem('token') || ''

  constructor() {
    makeAutoObservable(this)
  }

  setToken(token: string) {
    this.token = token
    localStorage.setItem('token', token)
  }
}

const authStore = new AuthStore()
export const useAuthStore = () => authStore