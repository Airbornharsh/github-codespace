import { useState } from 'react'

const Home = () => {
  const [gitLink, setGitLink] = useState('')
  const [name, setName] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const onClone = async () => {
    setIsLoading(true)
    try {
      const res = await fetch(
        `http://localhost:5000/api/git/clone?gitUrl=${gitLink}&name=${name}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: ''
        }
      )

      const data = await res.json()
      console.log(data)
      const projectName = gitLink.split('/').pop()?.replace('.git', '')
      if (!data.id) throw new Error('Failed to clone the repository')
      window.location.href = `http://${data.id}.localhost:5000?folder=/home/coder/Code/${projectName}`
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
          value={name}
          onChange={(e) => setName(e.target.value)}
          disabled={isLoading}
          placeholder="Your Name"
          className="h-10 bg-gray-300 px-2 outline-none rounded"
        />
        <input
          type="text"
          value={gitLink}
          onChange={(e) => setGitLink(e.target.value)}
          disabled={isLoading}
          placeholder="Git Link"
          className="h-10 bg-gray-300 px-2 outline-none rounded"
        />
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
