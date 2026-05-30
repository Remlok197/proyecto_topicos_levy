import React from 'react'
import HeaderPager from './HeaderPager'
import Form from './Form'

export default function CrearPage() {
  return (
    <main className="flex-1 overflow-y-auto bg-slate-50/50 p-6 md:p-8 flex flex-col items-center">
      <div className="w-full max-w-4xl">
        <HeaderPager />
        <Form />
      </div>
    </main>
  )
}
