import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useWebSocket } from '../context/WebSocketContext'
import FolderStructure from '../components/FolderStructure'
import CodeEditor from '@uiw/react-textarea-code-editor'
import { useFiles } from '../context/FilesContext'
import { BiSave, BiShareAlt } from 'react-icons/bi'
import Terminal from '../components/Terminal'

const Editor = () => {
  const { id } = useParams()
  const { activeFileData, activeFile } = useFiles()
  const { setSocket, saveFile } = useWebSocket()
  const [fileData, setFileData] = useState(activeFileData)
  const [timeo, setTimeo] = useState<NodeJS.Timeout | null>(null)
  const [isSaved, setIsSaved] = useState(false)

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
          <BiShareAlt
            color="white"
            size={'1.2rem'}
            className="cursor-pointer"
            onClick={() => {
              window.open(`http://${id}.localhost:5000/`, '_blank')
            }}
          />
          <BiSave
            color="white"
            size={'1.2rem'}
            className="cursor-pointer"
            onClick={() => {
              saveFile(activeFile, fileData)
              clearTimeout(timeo!)
              setIsSaved(true)
            }}
          />
        </div>
        <div>
          <p className="text-white px-2 text-xs">
            {activeFile}
            {isSaved ? '' : '*'}
          </p>
          <div className="overflow-auto h-[calc(100vh-19.5rem)]">
            <CodeEditor
              value={activeFileData}
              language={activeFile.split('.').pop()}
              placeholder="Choose a File"
              onChange={(e) => {
                setIsSaved(false)
                timeo && clearTimeout(timeo)
                setFileData(e.target.value)
                const t = setTimeout(() => {
                  saveFile(activeFile, e.target.value)
                  setIsSaved(true)
                }, 4000)
                setTimeo(t)
              }}
              padding={15}
              className="w-[calc(100vw-16rem)] text-white p-3 text-sm rounded-md outline-none overflow-scroll"
              style={{
                fontFamily:
                  'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace'
              }}
            />
          </div>
          <Terminal />
        </div>
      </div>
    </div>
  )
}

export default Editor
