
export function isEmpty<T>(value: T): value is Extract<T, undefined | null | ''> {
  return value === undefined || value === null || value === ''
}

export function omit<T extends object, K extends keyof T>(obj: T, keys: K[]): Omit<T, K> {
  const result = { ...obj }
  keys.forEach(key => {
    delete result[key]
  })
  return result as Omit<T, K>
}