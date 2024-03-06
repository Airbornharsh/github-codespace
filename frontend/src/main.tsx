import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { WebSocketProvider } from './context/WebSocketContext.tsx'
import { FilesProvider } from './context/FilesContext.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <FilesProvider>
    <WebSocketProvider>
      <App />
    </WebSocketProvider>
  </FilesProvider>
)
