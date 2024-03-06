import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useWebSocket } from '../context/WebSocketContext'

const Editor = () => {
  const { id } = useParams()
  const [dir, setDir] = useState<string>('/app')
  const [command, setCommand] = useState<string>('ls')
  const [type, setType] = useState<string>('files')
  const { setSocket, sendMessage } = useWebSocket()

  useEffect(() => {
    setSocket(id!)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id])

  return (
    <div>
      <input type="text" value={dir} onChange={(e) => setDir(e.target.value)} />
      <input
        type="text"
        value={command}
        onChange={(e) => setCommand(e.target.value)}
      />
      <input
        type="text"
        value={type}
        onChange={(e) => setType(e.target.value)}
      />
      <button
        onClick={() => {
          const cmd = {
            dir,
            command,
            type,
            isFile: 'README.md'
          }
          sendMessage(JSON.stringify(cmd))
        }}
      >
        Send
      </button>
    </div>
  )
}

export default Editor
