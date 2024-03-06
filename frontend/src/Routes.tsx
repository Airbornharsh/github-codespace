import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Editor from './pages/Editor'

const RoutesContainer = () => {
  return (
    <Routes>
      <Route path="/" Component={Home} />
      <Route path="/:id" Component={Editor} />
    </Routes>
  )
}

export default RoutesContainer
