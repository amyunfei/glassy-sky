import { useRef, useEffect } from 'react'

const useEventListener = (
  eventName: keyof GlobalEventHandlersEventMap,
  handler: EventListener,
  element: HTMLElement | Window & typeof globalThis | undefined = window
) => {
  const handlerRef = useRef<EventListener>(() => {})
  useEffect(() => {
    handlerRef.current = handler
  }, [handler])

  useEffect(() => {
    element.addEventListener(eventName, handlerRef.current)
    return () => {
      element.removeEventListener(eventName, handlerRef.current)
    }
  })
}
export default useEventListener