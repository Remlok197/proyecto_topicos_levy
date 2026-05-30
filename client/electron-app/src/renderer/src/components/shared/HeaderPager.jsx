import React from 'react'

export default function HeaderPager({ title, subtitle, rightElement }) {
  return (
    <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mb-6">
      <div>
        <h1 className="text-3xl font-bold text-neutral tracking-tight mb-1.5">{title}</h1>
        <p className="text-secondary text-sm md:text-base">{subtitle}</p>
      </div>
      {rightElement && <div className="self-start md:self-auto">{rightElement}</div>}
    </div>
  )
}
