import { useEffect } from 'react'
import { useFiles } from '../context/FilesContext'
import { useWebSocket } from '../context/WebSocketContext'
import Folder from './Folder'
import File from './File'

const FolderStructure = () => {
  const { files } = useFiles()
  const { sendMessage, socket } = useWebSocket()

  useEffect(() => {
    const cmd = {
      dir: '/app',
      command: 'ls',
      type: 'files',
      isFile: ''
    }
    setTimeout(() => {
      sendMessage(JSON.stringify(cmd))
    }, 200)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [socket])

  return (
    <div className="bg-fStructBackground w-64 h-screen fixed left-0 top-0 z-30 overflow-auto">
      <div className="p-2">
        <h1 className="text-white text-xl">Folder Structure</h1>
        <ul className="text-sm flex flex-col gap-1">
          {Object.keys(files).map((file) => {
            const data = files[file].path.split('/')
            if (data.length > 1) {
              return null
            }
            if (files[file].type === 'file') {
              return <File key={file} path={files[file].path} />
            } else {
              return (
                <Folder
                  key={file}
                  path={files[file].path.split('/app')[1] || files[file].path}
                />
              )
            }
          })}
        </ul>
      </div>
    </div>
  )
}

export default FolderStructure
