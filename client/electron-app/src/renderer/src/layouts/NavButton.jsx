import { NavLink } from 'react-router'

export default function NavButton({ to, icon, children }) {
  return (
    <NavLink
      to={to}
      className={({ isActive }) =>
        `flex items-center gap-2.5 px-4 py-2 rounded-lg text-sm font-mediu ansition-all duration-200 select-none ${
          isActive
            ? 'bg-primary/15 text-primary font-semibold'
            : 'text-secondary hover:bg-tertiary hover:text-neutral'
        }`
      }
    >
      {icon && <span className="flex items-center justify-center shrink-0">{icon}</span>}
      <span className="leading-none">{children}</span>
    </NavLink>
  )
}
