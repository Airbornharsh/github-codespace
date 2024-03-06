import React, { createContext, useContext, useState, ReactNode } from 'react'
import { useFiles } from './FilesContext'

interface WebSocketContextProps {
  socket: WebSocket | null
  setSocket: (id: string) => void
  message: string
  sendMessage: (path: string) => void
  getFile: (path: string) => void
  saveFile: (path: string, data: string) => void
}

const WebSocketContext = createContext<WebSocketContextProps | undefined>(
  undefined
)

// eslint-disable-next-line react-refresh/only-export-components
export const useWebSocket = () => {
  const context = useContext(WebSocketContext)

  if (!context) {
    throw new Error('useWebSocket must be used within a WebSocketProvider')
  }

  return context
}

interface WebSocketProviderProps {
  children: ReactNode
}

export const WebSocketProvider: React.FC<WebSocketProviderProps> = ({
  children
}) => {
  const [socket, setSocket] = useState<WebSocket | null>(null)
  const [message, setMessage] = useState<string>('')
  const { setFiles, setFileData } = useFiles()

  const setSocketFn = (id: string) => {
    if (socket) {
      socket.close()
    }

    const newSocket = new WebSocket(`ws://${id}.localhost:5000/ws`)

    newSocket.addEventListener('open', (event) => {
      console.log('WebSocket connection opened:', event)
    })

    newSocket.addEventListener('close', (event) => {
      console.log('WebSocket connection closed:', event)
    })

    newSocket.addEventListener('error', (event) => {
      console.error('WebSocket connection error:', event)
    })

    newSocket.addEventListener('message', (event) => {
      const resData = JSON.parse(event.data)
      if (resData.type === 'files') {
        setFiles(
          resData.dir.split('/app/')[1] || '',
          resData.out.split('\n').filter((t: string) => t !== '')
        )
      } else if (resData.type === 'file') {
        console.log(resData)
        setFileData(
          resData.dir.split('/app/')[1] + '/' + resData.isFile ||
            '' + resData.isFile,
          resData.out
        )
      } else if (resData.type === 'command') {
        setMessage(resData.out)
      }
    })

    setSocket(newSocket)
  }

  const sendMessage = (message: string) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(message)
    }
  }

  const getFile = (path: string) => {
    const filteredPath = path.split('/').filter((e) => e !== '')
    const removedFile = filteredPath.slice(0, -1).join('/')
    const fileName = filteredPath.pop()
    const cmd = {
      dir: '/app/' + removedFile,
      command: 'cat ' + fileName,
      type: 'file',
      isFile: fileName
    }
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(cmd))
    }
  }

  const saveFile = (path: string, data: string) => {
    console.log(path)
    const filteredPath = path.split('/').filter((e) => e !== '')
    const removedFile = filteredPath.slice(0, -1).join('/')
    const fileName = filteredPath.pop()
    const cmd = {
      dir: '/app/' + removedFile,
      command: `echo "` + data + `" > ` + fileName,
      type: 'command',
      isFile: fileName
    }
    console.log(cmd)
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(cmd))
    }
  }

  const contextValue: WebSocketContextProps = {
    socket,
    setSocket: setSocketFn,
    message,
    sendMessage,
    getFile,
    saveFile
  }

  return (
    <WebSocketContext.Provider value={contextValue}>
      {children}
    </WebSocketContext.Provider>
  )
}
