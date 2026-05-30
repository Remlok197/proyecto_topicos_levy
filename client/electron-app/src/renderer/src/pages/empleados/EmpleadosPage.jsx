import React, { useState } from 'react'
import HeaderPager from '../../components/shared/HeaderPager'
import EmpleadosTable from './EmpleadosTable'

export default function EmpleadosPage() {
  const [activeTab, setActiveTab] = useState('activos')

  const tabsElement = (
    <div className="flex bg-slate-100 p-1 rounded-lg border border-slate-200">
      <button
        onClick={() => setActiveTab('activos')}
        className={`px-4 py-1.5 rounded-md text-xs font-semibold transition-all cursor-pointer ${
          activeTab === 'activos'
            ? 'bg-white text-neutral shadow-sm'
            : 'text-secondary hover:text-neutral'
        }`}
      >
        Empleados Activos
      </button>
      <button
        onClick={() => setActiveTab('ex')}
        className={`px-4 py-1.5 rounded-md text-xs font-semibold transition-all cursor-pointer ${
          activeTab === 'ex'
            ? 'bg-white text-neutral shadow-sm'
            : 'text-secondary hover:text-neutral'
        }`}
      >
        Ex-empleados
      </button>
    </div>
  )

  return (
    <main className="flex-1 overflow-y-auto bg-slate-50/50 p-6 md:p-8 flex flex-col items-center">
      <div className="w-full max-w-6xl">
        <HeaderPager
          title="Directorio de Empleados"
          subtitle="Consulte la lista de empleados activos y realice modificaciones rápidas en sus perfiles."
          rightElement={tabsElement}
        />
        <EmpleadosTable />
      </div>
    </main>
  )
}
