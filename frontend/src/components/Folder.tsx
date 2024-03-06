import { useState } from 'react'
import { IoMdArrowDropdown, IoMdArrowDropup } from 'react-icons/io'
import { useWebSocket } from '../context/WebSocketContext'
import { useFiles } from '../context/FilesContext'
import File from './File'

interface FolderProps {
  path: string
}

const Folder: React.FC<FolderProps> = ({ path }) => {
  const [isOpen, setIsOpen] = useState(false)
  const { files } = useFiles()
  const { sendMessage } = useWebSocket()

  return (
    <li className="text-white flex flex-col">
      <div
        className="flex justify-between cursor-pointer"
        onClick={() => {
          setIsOpen((o) => !o)
          const cmd = {
            dir: '/app/' + path,
            command: 'ls',
            type: 'files',
            isFile: ''
          }
          sendMessage(JSON.stringify(cmd))
        }}
      >
        <p>
          {path
            .split('/')
            .filter((e) => e !== '')
            .pop()}
        </p>
        {isOpen ? <IoMdArrowDropdown /> : <IoMdArrowDropup />}
      </div>
      {isOpen && (
        <ul className="ml-2">
          {Object.keys(files).map((file) => {
            if (!files[file].path) {
              return null
            }
            if (
              files[file].path.includes(path) &&
              files[file].path.split('/').filter((e) => e !== '').length ===
                path.split('/').filter((e) => e !== '').length + 1
            ) {
              if (files[file].type === 'file') {
                return <File key={file} path={file} />
              } else {
                return (
                  <Folder
                    key={file}
                    path={files[file].path.split('/app')[1] || files[file].path}
                  />
                )
              }
            }
            return null
          })}
        </ul>
      )}
    </li>
  )
}

export default Folder
