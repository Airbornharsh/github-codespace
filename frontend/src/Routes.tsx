import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import VsCode from './pages/VsCode'

const RoutesContainer = () => {
  return (
    <Routes>
      <Route path="/" Component={Home} />
      <Route path="/:id" Component={VsCode} />
    </Routes>
  )
}

export default RoutesContainer
