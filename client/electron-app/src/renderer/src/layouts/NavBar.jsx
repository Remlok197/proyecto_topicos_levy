import { Calendar, Users, Columns3, UserPlus, UserMinus } from "lucide-react";
import NavButton from "./NavButton";

export default function NavBar() {
  const navItems = [
    { to: "/historial", label: "Historial", icon: <Calendar size={18} /> },
    { to: "/empleados", label: "Empleados", icon: <Users size={18} /> },
    { to: "/transferir", label: "Transferir", icon: <Columns3 size={18} /> },
    { to: "/crear", label: "Crear", icon: <UserPlus size={18} /> },
    { to: "/borrar", label: "Borrar", icon: <UserMinus size={18} /> },
  ];

  return (
    <nav className="flex items-center gap-2">
      {navItems.map((item) => (
        <NavButton key={item.to} to={item.to} icon={item.icon}>
          {item.label}
        </NavButton>
      ))}
    </nav>
  );
}