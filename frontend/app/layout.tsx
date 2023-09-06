import './globals.css'
import type { Metadata } from 'next'
import SideBar from '../components/sidebar'


export const metadata: Metadata = {
  title: 'Allegro',
  description: 'A distributed music streaming service for classicals',
}

export default function RootLayout({children}: {children: React.ReactNode}) {
  return (
  <>
  <SideBar />
  <main>{children}</main>
  </>
  )
}
