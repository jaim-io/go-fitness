export interface IAPIResult<T>{
  result: T | undefined,
  error: Error | undefined
}