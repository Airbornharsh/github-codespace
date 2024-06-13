import { useEffect, useState } from 'react'
import { useFiles } from '../context/FilesContext'
import { useWebSocket } from '../context/WebSocketContext'
import Folder from './Folder'
import File from './File'
import { BiPlus } from 'react-icons/bi'
import { CgClose } from 'react-icons/cg'

const FolderStructure = () => {
  const [addClicked, setAddClicked] = useState(false)
  const [addHover, setAddHover] = useState(false)
  const [addType, setAddType] = useState<'folder' | 'file'>('file')
  const [addName, setAddName] = useState('')
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

  const createFileOrFolder = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (
      addName === '' ||
      addName.includes('/') ||
      addName.includes(' ') ||
      addName.includes('.')
    ) {
      return
    }

    const cmd = {
      dir: '/app',
      command:
        addType === 'file'
          ? 'touch ' + addName + ' && echo ' + addName
          : 'mkdir ' + addName + ' && echo ' + addName,
      type: 'files',
      isFile: addType
    }
    sendMessage(JSON.stringify(cmd))
    setAddClicked(false)
    setAddName('')
  }

  return (
    <div className="bg-fStructBackground w-64 h-screen fixed left-0 top-0 z-30 overflow-auto border-r-[0.1rem] border-gray-600">
      <div className="p-2">
        <div
          className="flex justify-between items-center"
          onMouseOver={() => {
            setAddHover(true)
          }}
          onMouseLeave={() => {
            setAddHover(false)
          }}
        >
          <h1 className="text-white text-lg">Folder Structure</h1>
          <BiPlus
            className="text-white text-lg cursor-pointer"
            onClick={() => {
              setAddClicked((a) => !a)
            }}
          />
          {addHover && (
            <span className="absolute right-0 top-8 bg-gray-300 flex flex-col w-24 items-center z-50">
              <p
                className="hover:bg-gray-600 hover:text-white w-24 flex items-center justify-center cursor-pointer"
                onClick={() => {
                  setAddType('folder')
                  setAddClicked(true)
                }}
              >
                Folder
              </p>
              <p
                className="hover:bg-gray-600 hover:text-white w-24 flex items-center justify-center cursor-pointer"
                onClick={() => {
                  setAddType('file')
                  setAddClicked(true)
                }}
              >
                File
              </p>
            </span>
          )}
        </div>

        <ul className="text-sm flex flex-col gap-1">
          {addClicked && (
            <form
              className="border-[0.1rem] relative text-white text-sm"
              onSubmit={createFileOrFolder}
            >
              <input
                type="text"
                className="w-full bg-transparent outline-none pl-1"
                value={addName}
                onChange={(e) => {
                  setAddName(e.target.value)
                }}
              />
              <p className="absolute right-[1.2rem] top-[0.02rem] text-xs">
                *{addType}
              </p>
              <CgClose
                className="absolute right-0 top-0 text-white text-lg cursor-pointer"
                onClick={() => {
                  setAddClicked(false)
                }}
              />
            </form>
          )}
          {Object.keys(files).map((file) => {
            if (!files[file].path) {
              return null
            }
            const data = files[file].path.split('/')
            if (data.length > 1) {
              return null
            }
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
          })}
        </ul>
      </div>
    </div>
  )
}

export default FolderStructure
