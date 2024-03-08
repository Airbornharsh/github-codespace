import { useEffect, useRef } from 'react'
import { useTerminal } from '../context/TerminalContext'
import { useWebSocket } from '../context/WebSocketContext'

const Terminal = () => {
  const { outputs, routes, activeCommand, setActiveCommand } =
    useTerminal()
  const { execCommand } = useWebSocket()
  const terminalRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLInputElement>(null)

  const execCmd = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    execCommand()
    console.log('executing command:', activeCommand)
  }

  const scrollToBottom = () => {
    if (terminalRef.current) {
      terminalRef.current.scrollTop = terminalRef.current.scrollHeight
    }
  }

  useEffect(() => {
    scrollToBottom()
  }, [outputs])

  return (
    <div
      className="h-64 bg-fStructBackground p-2 overflow-auto"
      id="terminal"
      ref={terminalRef}
      onClick={() => {
        inputRef.current?.focus()
      }}
    >
      <ul>
        {outputs.map((out, i) => {
          return (
            <li key={out.command + out.dir + i}>
              <div className="flex text-sm gap-2 items-center">
                <span className="text-blue-600 flex">
                  <pre className="flex">
                    {out.oldDir
                      .split('/')
                      .filter((o) => o !== '')
                      .join('/')}
                  </pre>
                  <pre>/</pre>
                </span>
                <pre
                  key={out.command}
                  className="bg-transparent outline-none text-gray-500 w-full"
                >
                  {out.command}
                </pre>
              </div>
              <pre className="text-gray-500">{out.out}</pre>
            </li>
          )
        })}
      </ul>
      <form
        className="flex text-sm gap-2 items-center"
        onSubmit={execCmd}
        id="terminal-input"
      >
        <span className="text-blue-600 flex">
          <pre>{routes.join('/')}</pre>
          <pre>/</pre>
        </span>
        <pre>
          <input
            type="text"
            ref={inputRef}
            className="bg-transparent outline-none text-gray-500 w-full"
            value={activeCommand}
            onChange={(e) => setActiveCommand(e.target.value)}
          />
        </pre>
      </form>
    </div>
  )
}

export default Terminal
