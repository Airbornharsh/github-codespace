import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { WebSocketProvider } from './context/WebSocketContext.tsx'
import { FilesProvider } from './context/FilesContext.tsx'
import { TerminalProvider } from './context/TerminalContext.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <TerminalProvider>
    <FilesProvider>
      <WebSocketProvider>
        <App />
      </WebSocketProvider>
    </FilesProvider>
  </TerminalProvider>
)
