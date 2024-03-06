import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useWebSocket } from '../context/WebSocketContext'
import FolderStructure from '../components/FolderStructure'
import CodeEditor from '@uiw/react-textarea-code-editor'
import { useFiles } from '../context/FilesContext'
import { BiSave } from 'react-icons/bi'

const Editor = () => {
  const { id } = useParams()
  const { activeFileData, activeFile } = useFiles()
  const { setSocket, saveFile } = useWebSocket()
  const [fileData, setFileData] = useState(activeFileData)
  const [timeo, setTimeo] = useState<NodeJS.Timeout | null>(null)

  useEffect(() => {
    setSocket(id!)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id])

  return (
    <div className="w-screen h-screen bg-background flex">
      <FolderStructure />
      <div className="w-[calc(100vw-16rem)] max-h-screen ml-64">
        <div className="bg-fStructBackground w-[100%] flex justify-between p-2 pr-3 h-10">
          <p></p>
          <BiSave
            color="white"
            size={'1.2rem'}
            className="cursor-pointer"
            onClick={() => {
              saveFile(activeFile, fileData)
            }}
          />
        </div>
        <div className="overflow-auto h-[100vh-2.5rem]">
          <CodeEditor
            value={activeFileData}
            language={activeFile.split('.').pop()}
            placeholder="Choose a File"
            disabled={!activeFileData}
            onChange={(e) => {
              timeo && clearTimeout(timeo)
              setFileData(e.target.value)
              const t = setTimeout(() => {
                saveFile(activeFile, e.target.value)
              }, 2000)
              setTimeo(t)
            }}
            padding={15}
            className="w-[calc(100vw-16rem)] text-white p-3 text-sm rounded-md outline-none"
            style={{
              fontFamily:
                'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace'
            }}
          />
        </div>
      </div>
    </div>
  )
}

export default Editor
