import { createContext, useContext, useState } from 'react'

interface FileType {
  type: string
  path: string
  data?: string
}

interface FolderType {
  type: string
  path: string
}

interface FilesContextProps {
  files: {
    [key: string]: FileType | FolderType
  }
  setFiles: (dir: string, tempFiles: string[]) => void
  setFileData: (dir: string, data: string) => void
  activeFile: string
  activeFileData: string
}

const FilesContext = createContext<FilesContextProps | undefined>(undefined)

// eslint-disable-next-line react-refresh/only-export-components
export const useFiles = () => {
  const context = useContext(FilesContext)

  if (!context) {
    throw new Error('useFiles must be used within a FilesProvider')
  }

  return context
}

interface FilesProviderProps {
  children: React.ReactNode
}

export const FilesProvider: React.FC<FilesProviderProps> = ({ children }) => {
  const [files, setFiles] = useState<{
    [key: string]: FileType | FolderType
  }>({})
  const [activeFile, setActiveFile] = useState('')
  const [activeFileData, setActiveFileData] = useState('')

  const setFilesFn = (dir: string, tempFileNames: string[]) => {
    tempFileNames.forEach((file) => {
      const fileData = {
        type: file.includes('.') ? 'file' : 'folder',
        path: `${dir ? dir + '/' : ''}${file}`,
        data: ''
      }
      setFiles((prevFiles) => {
        return {
          ...prevFiles,
          [fileData.path]: fileData
        }
      })
    })
  }

  const setFileData = (dir: string, data: string) => {
    setActiveFile(dir)
    setActiveFileData(data)
    setFiles((prevFiles) => {
      return {
        ...prevFiles,
        [dir]: {
          ...prevFiles[dir],
          data
        }
      }
    })
  }

  console.log(files)

  return (
    <FilesContext.Provider
      value={{
        files,
        setFiles: setFilesFn,
        setFileData,
        activeFile,
        activeFileData
      }}
    >
      {children}
    </FilesContext.Provider>
  )
}
