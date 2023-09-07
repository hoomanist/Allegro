import Image from 'next/image'


export default async function Card({childern, imgSrc, ...props}) {
    return (
        <div {...props} className="">
            <Image src={imgSrc} alt="" width={450} height={250}
                        className="rounded-t-xl flex w-full aspect-square"/>
            <div className="">
                <div className="p-4 text-white">{childern}</div>
            </div>
        </div>
    )
}
