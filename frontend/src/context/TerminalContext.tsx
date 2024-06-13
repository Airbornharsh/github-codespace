import { createContext, useContext, useState } from 'react'

interface CommandsType {
  command: string
  oldDir: string
  dir: string
  out: string
}

interface TerminalContextProps {
  routes: string[]
  setRoutes: (routes: string[]) => void
  outputs: CommandsType[]
  setOutputs: (output: CommandsType, clear: boolean) => void
  activeCommand: string
  setActiveCommand: (command: string) => void
}

const TerminalContext = createContext<TerminalContextProps | undefined>(
  undefined
)

// eslint-disable-next-line react-refresh/only-export-components
export const useTerminal = () => {
  const context = useContext(TerminalContext)

  if (!context) {
    throw new Error('useFiles must be used within a FilesProvider')
  }

  return context
}

interface TerminalProviderProps {
  children: React.ReactNode
}

export const TerminalProvider: React.FC<TerminalProviderProps> = ({
  children
}) => {
  const [routes, setRoutes] = useState<string[]>(['app'])
  const [outputs, setOutputs] = useState<CommandsType[]>([])
  const [activeCommand, setActiveCommand] = useState<string>('')

  const setRoutesFn = (routes: string[]) => {
    setRoutes([...routes])
  }

  const setOutputsFn = (output: CommandsType, empty = false) => {
    setActiveCommand('')
    if (empty) {
      setOutputs([])
      return
    }
    setOutputs((o) => [...o, output])
    setRoutesFn(
      output.dir
        .replace('\n', '')
        .split('/')
        .filter((o) => o !== '')
    )
  }

  return (
    <TerminalContext.Provider
      value={{
        routes,
        setRoutes: setRoutesFn,
        outputs,
        setOutputs: setOutputsFn,
        activeCommand,
        setActiveCommand,
      }}
    >
      {children}
    </TerminalContext.Provider>
  )
}
