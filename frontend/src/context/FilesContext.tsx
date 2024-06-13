import { createContext, useContext, useEffect, useState } from 'react'

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
  removeFile: (dir: string) => void
  removeFolder: (dir: string) => void
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

  useEffect(() => {
    const compareFiles = (a: string, b: string): number => {
      const typeA = files[a].type
      const typeB = files[b].type

      if (typeA === 'folder' && typeB !== 'folder') {
        return -1
      } else if (typeA !== 'folder' && typeB === 'folder') {
        return 1
      } else {
        return a.localeCompare(b)
      }
    }

    const sortedFileKeys = Object.keys(files).sort(compareFiles)

    const sortedFiles = sortedFileKeys.reduce(
      (acc, key) => {
        acc[key] = files[key]
        return acc
      },
      {} as { [key: string]: FileType | FolderType }
    )

    setFiles(sortedFiles)
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Object.keys(files).length])

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

  const removeFile = (dir: string) => {
    setFiles((prevFiles) => {
      const newFiles = { ...prevFiles }
      delete newFiles[dir]
      return newFiles
    })
  }

  const removeFolder = (dir: string) => {
    setFiles((prevFiles) => {
      const newFiles = { ...prevFiles }
      Object.keys(newFiles).forEach((file) => {
        if (file.includes(dir)) {
          delete newFiles[file]
        }
      })
      return newFiles
    })
  }

  return (
    <FilesContext.Provider
      value={{
        files,
        setFiles: setFilesFn,
        setFileData,
        activeFile,
        activeFileData,
        removeFile,
        removeFolder
      }}
    >
      {children}
    </FilesContext.Provider>
  )
}
