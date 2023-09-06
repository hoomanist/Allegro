async function getData() {
  const res = await fetch('http://127.0.0.1:8001/api/ping')
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json()
}
 

export default async function Home() {
    const data = await getData()
    
    return (
        <div className="grid h-screen place-items-center ">
        Hi
        </div>
    )
}
