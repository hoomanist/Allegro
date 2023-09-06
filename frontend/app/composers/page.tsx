import Card from '../../components/card'
import {FiMic} from "@react-icons/all-files/fi/FiMic"
async function getData() {
    const res = await fetch("http://127.0.0.1:8001/api/q/composers")
    if (!res.ok) {
        throw new Error("unable to fetch data")
    }
    return res.json()
}

export default async function Composer() {

    const imageUrl = "http://127.0.0.1:8001/static/"
    const data = await getData()

    console.log(data[0]["name"])
    return (
        <div className="grid grid-cols-4 ">
            {data.map(item => (
            <Card imgSrc={imageUrl + item["photo"]}>
            <div >
                <h3 className='text-xl font-bold mb-2'>{item["name"]}</h3>
                <p>
                    {item["description"]}
                </p>
            </div>
            </Card>
            ))}
        </div>
    )
}
