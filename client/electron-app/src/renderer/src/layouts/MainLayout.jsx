import { Outlet } from 'react-router'
import Header from './Header'

export default function MainLayout() {
  return (
    <div className="flex flex-col w-screen h-screen">
      <Header />
      <Outlet />
    </div>
  )
}
