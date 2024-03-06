import { useFiles } from '../context/FilesContext'
import { useWebSocket } from '../context/WebSocketContext'

interface FileProps {
  path: string
}

const File: React.FC<FileProps> = ({ path }) => {
  const { getFile } = useWebSocket()
  const { activeFile } = useFiles()

  return (
    <li
      className={`text-white flex justify-between cursor-pointer ${activeFile === path ? 'bg-gray-800' : 'bg-transparent'}`}
      onClick={() => {
        console.log(path)
        getFile(path)
      }}
    >
      {path
        .split('/')
        .filter((e) => e !== '')
        .pop()}
    </li>
  )
}

export default File
