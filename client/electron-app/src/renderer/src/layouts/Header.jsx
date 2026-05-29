import NavBar from "./NavBar";

export default function Header() {
  return (
    <header className="flex items-center justify-between px-8 py-4 bg-white border-b border-slate-100 shadow-sm">
      <div className="text-lg font-semibold text-slate-800 tracking-tight">
        Administración de Empleados
      </div>
      <NavBar />
    </header>
  );
}