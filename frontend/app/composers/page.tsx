import Card from '../../components/card'
async function getData() {
    const res = await fetch("http://127.0.0.1:8001/api/q/composers", {cache: 'no-cache'})
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
        <div className="grid grid-cols-5 gap-x-3 gap-y-2 ml-20 mt-10 mr-5 z-[-1] justify-center">
            {data.map(item => (
            <div className="scale-[0.9] rounded-xl shadown-lg border-black bg-gradient-to-t from-zinc-200 to-amber-100 group">
            <Card imgSrc={imageUrl + item["photo"]}>
            </Card>
            <div className="m-2 items-center ">
                <h3 className="text-xl font-semibold mb-2">{item["name"]}</h3>
                <h4 className="text-sm font-mono text-zinc-600">{item["birth"]} - {item["death"]}</h4>
                <p>
                    {item["desc"]}
                </p>
            </div>
            </div>
            ))}

        </div>
    )
}
