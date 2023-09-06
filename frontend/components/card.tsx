export default async function Card({childern, imgSrc, ...props}) {
    return (
        <div {...props} className="relative max-w-xs overflow-hidden rounded-2xl 
        shadow-lg scale-[0.5] w-full group">
            <img src={imgSrc} alt="" className="transition-transform 
            group-hover:scale-110 duration-200 object-fill"/>
            <div className="absolute inset-0 flex items-end 
            bg-gradient-to-t from-black/60 to-transparent">
                <div className="p-4 text-white">{childern}</div>
            </div>
        </div>
    )
}
