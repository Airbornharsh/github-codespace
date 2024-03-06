import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Home = () => {
  const [gitLink, setGitLink] = useState('')
  const [rootDir, setRootDir] = useState('')
  const [stack, setStack] = useState('nextjs')
  const [isLoading, setIsLoading] = useState(false)
  const Navigate = useNavigate()

  const onClone = async () => {
    setIsLoading(true)
    try {
      const res = await fetch(
        `http://localhost:5000/api/git/clone?git-link=${gitLink}&root-dir=${rootDir}&stack=${stack}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: ''
        }
      )

      const data = await res.json()
      Navigate(`/${data.repo}`, { replace: false })
    } catch (e) {
      console.error(e)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="w-screen h-screen flex justify-center items-center flex-col gap-2">
      <div className="flex flex-col w-1/2 max-w-[30rem] gap-2">
        <input
          type="text"
          value={gitLink}
          onChange={(e) => setGitLink(e.target.value)}
          disabled={isLoading}
          placeholder="Git Link"
          className="h-10 bg-gray-300 px-2 outline-none rounded"
        />
        <input
          type="text"
          value={rootDir}
          onChange={(e) => setRootDir(e.target.value)}
          disabled={isLoading}
          placeholder="Root Directory"
          className="h-10 bg-gray-300 px-2 outline-none rounded"
        />
        <select
          value={stack}
          onChange={(e) => setStack(e.target.value)}
          disabled={isLoading}
          className="h-10 bg-gray-300 px-2 outline-none rounded"
        >
          <option value="nextjs">Next.js</option>
          <option value="reactjs">React</option>
        </select>
      </div>
      <button
        className="bg-blue-500 text-white px-4 py-2 rounded-md ml-2"
        disabled={isLoading}
        onClick={onClone}
      >
        {isLoading ? 'Cloning...' : 'Clone'}
      </button>
    </div>
  )
}

export default Home
