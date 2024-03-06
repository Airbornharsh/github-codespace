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

// const data = {
//   type: 'folder',
//   path: 'app',
//   children: {
//     'src': {
//       type: 'folder',
//       path: 'app/src',
//       children: {
//         'index.js': {
//           type: 'file',
//           path: 'app/src/index.js',
//           data: 'console.log("Hello, world!")'
//         }
//       }
//     },
//     'index.html': {
//       type: 'file',
//       path: 'app/index.html',
//       data: '<h1>Hello, world!</h1>'
//     }
//   }
// }

interface FilesContextProps {
  files: {
    [key: string]: FileType | FolderType
  }
  setFiles: (dir: string, tempFiles: string[]) => void
  setFileData: (dir: string, data: string) => void
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
    <FilesContext.Provider value={{ files, setFiles: setFilesFn, setFileData }}>
      {children}
    </FilesContext.Provider>
  )
}
