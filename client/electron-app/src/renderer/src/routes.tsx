import { createBrowserRouter } from 'react-router'
import MainLayout from './layouts/MainLayout'
import EmpleadosPage from './pages/empleados/EmpleadosPage'
import HistorialPage from './pages/historial/HistorialPage'
import TransferirPage from './pages/transferir/TransferirPage'
import CrearPage from './pages/crear/CrearPage'
import BorrarPage from './pages/borrar/BorrarPage'

export const router = createBrowserRouter([
  {
    path: '/',
    Component: MainLayout,
    children: [
      { path: 'historial', element: <HistorialPage /> },
      { path: 'empleados', element: <EmpleadosPage /> },
      { path: 'transferir', element: <TransferirPage /> },
      { path: 'crear', element: <CrearPage /> },
      { path: 'borrar', element: <BorrarPage /> }
    ]
  }
])
