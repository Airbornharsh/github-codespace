import { useParams } from 'react-router-dom'

const VsCode = () => {
  const { id } = useParams()
  return (
    <div>
      <iframe
        src={`http://${id}.localhost:5000`}
        title="VsCode"
        className="w-screen h-screen"
      ></iframe>
    </div>
  )
}

export default VsCode
