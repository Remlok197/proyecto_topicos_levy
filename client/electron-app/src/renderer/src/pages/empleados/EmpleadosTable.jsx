import React, { useState } from 'react'
import { Search, Mail, Pencil, Check, X, ChevronLeft, ChevronRight, User } from 'lucide-react'

export default function EmpleadosTable() {
  // Mock data representing database employees
  const [employees, setEmployees] = useState([
    {
      employeeNumber: 8492,
      firstName: 'Elena',
      lastName: 'Rodríguez',
      department: 'IT',
      jobTitle: 'Developer',
      email: 'e.rodriguez@corp.com',
      avatarColor: 'bg-teal-100 text-teal-700'
    },
    {
      employeeNumber: 7310,
      firstName: 'Marcus',
      lastName: 'Chen',
      department: 'IT',
      jobTitle: 'Manager',
      email: 'm.chen@corp.com',
      avatarColor: 'bg-blue-100 text-blue-700'
    },
    {
      employeeNumber: 9021,
      firstName: 'Sarah',
      lastName: 'Williams',
      department: 'HR',
      jobTitle: 'Specialist',
      email: 's.williams@corp.com',
      avatarColor: 'bg-purple-100 text-purple-700'
    },
    {
      employeeNumber: 4219,
      firstName: 'Daniel',
      lastName: 'Gómez',
      department: 'Sales',
      jobTitle: 'Analyst',
      email: 'd.gomez@corp.com',
      avatarColor: 'bg-amber-100 text-amber-700'
    },
    {
      employeeNumber: 5543,
      firstName: 'Laura',
      lastName: 'Martínez',
      department: 'Finance',
      jobTitle: 'Specialist',
      email: 'l.martinez@corp.com',
      avatarColor: 'bg-rose-100 text-rose-700'
    }
  ])

  // Search filter state
  const [searchTerm, setSearchTerm] = useState('')

  // Inline editing state
  const [editingId, setEditingId] = useState(null)
  const [editFormData, setEditFormData] = useState({
    firstName: '',
    lastName: '',
    department: '',
    jobTitle: '',
    email: ''
  })

  // Validation errors
  const [validationErrors, setValidationErrors] = useState({})

  // Departments list
  const departments = [
    { value: 'HR', label: 'Recursos Humanos' },
    { value: 'IT', label: 'Tecnología' },
    { value: 'Sales', label: 'Ventas' },
    { value: 'Finance', label: 'Finanzas' },
    { value: 'Marketing', label: 'Mercadotecnia' }
  ]

  // Roles list
  const roles = [
    { value: 'Developer', label: 'Desarrollador' },
    { value: 'Designer', label: 'Diseñador' },
    { value: 'Manager', label: 'Gerente' },
    { value: 'Analyst', label: 'Analista' },
    { value: 'Specialist', label: 'Especialista' }
  ]

  // Start editing a row
  const handleEditStart = (employee) => {
    setEditingId(employee.employeeNumber)
    setEditFormData({
      firstName: employee.firstName,
      lastName: employee.lastName,
      department: employee.department,
      jobTitle: employee.jobTitle,
      email: employee.email
    })
    setValidationErrors({})
  }

  // Change values in edit form
  const handleEditChange = (e) => {
    const { name, value } = e.target
    setEditFormData((prev) => ({
      ...prev,
      [name]: value
    }))

    if (validationErrors[name]) {
      setValidationErrors((prev) => ({
        ...prev,
        [name]: null
      }))
    }
  }

  // Save inline edit changes
  const handleEditSave = (id) => {
    // Validate fields
    const errors = {}
    if (!editFormData.firstName.trim()) errors.firstName = true
    if (!editFormData.lastName.trim()) errors.lastName = true
    if (!editFormData.department) errors.department = true
    if (!editFormData.jobTitle) errors.jobTitle = true
    if (!editFormData.email.trim()) {
      errors.email = true
    } else if (!/\S+@\S+\.\S+/.test(editFormData.email)) {
      errors.email = true
    }

    if (Object.keys(errors).length > 0) {
      setValidationErrors(errors)
      return
    }

    // Update employees state
    setEmployees((prev) =>
      prev.map((emp) =>
        emp.employeeNumber === id
          ? {
              ...emp,
              firstName: editFormData.firstName.trim(),
              lastName: editFormData.lastName.trim(),
              department: editFormData.department,
              jobTitle: editFormData.jobTitle,
              email: editFormData.email.trim()
            }
          : emp
      )
    )

    // Reset editing state
    setEditingId(null)
  }

  // Cancel inline editing
  const handleEditCancel = () => {
    setEditingId(null)
    setValidationErrors({})
  }

  // Filter employees list
  const filteredEmployees = employees.filter((emp) => {
    const fullName = `${emp.firstName} ${emp.lastName}`.toLowerCase()
    const deptLabel = (
      departments.find((d) => d.value === emp.department)?.label || ''
    ).toLowerCase()
    const roleLabel = (roles.find((r) => r.value === emp.jobTitle)?.label || '').toLowerCase()
    const email = emp.email.toLowerCase()
    const search = searchTerm.toLowerCase()

    return (
      fullName.includes(search) ||
      deptLabel.includes(search) ||
      roleLabel.includes(search) ||
      email.includes(search)
    )
  })

  const getInitials = (first, last) => {
    return `${first.charAt(0)}${last.charAt(0)}`.toUpperCase()
  }

  // Input styles
  const inputClass =
    'bg-white border text-xs text-neutral rounded px-2 py-1 transition-all focus:outline-none focus:ring-1 focus:ring-primary/30 focus:border-primary'
  const errorInputClass =
    'bg-white border border-red-500 text-xs text-neutral rounded px-2 py-1 focus:outline-none focus:ring-1 focus:ring-red-200 focus:border-red-500'

  return (
    <div className="bg-white border border-slate-200 rounded-xl shadow-sm overflow-hidden w-full flex flex-col">
      {/* Search Header */}
      <div className="p-4 bg-slate-50 border-b border-slate-200 flex items-center justify-between gap-4">
        <div className="relative w-full max-w-sm">
          <Search size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
          <input
            type="text"
            placeholder="Buscar por nombre, correo, departamento..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="w-full bg-white border border-slate-200 text-xs text-neutral placeholder-slate-400 rounded-lg pl-9 pr-4 py-2 focus:outline-none focus:border-primary/50 focus:ring-2 focus:ring-primary/10 transition-all"
          />
        </div>
        <div className="text-xs text-secondary font-medium">
          {filteredEmployees.length} resultado{filteredEmployees.length !== 1 ? 's' : ''}
        </div>
      </div>

      {/* Table Container */}
      <div className="overflow-x-auto">
        <table className="w-full text-left border-collapse">
          <thead>
            <tr className="bg-slate-50 border-b border-slate-200 text-[10px] font-bold text-secondary uppercase tracking-wider">
              <th className="px-6 py-3.5 w-1/4">Nombre del Empleado</th>
              <th className="px-6 py-3.5 w-1/4">Puesto y Departamento</th>
              <th className="px-6 py-3.5 w-1/4">Contacto</th>
              <th className="px-6 py-3.5 w-1/8">Estado</th>
              <th className="px-6 py-3.5 w-1/8 text-right">Acciones</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100 text-sm text-neutral">
            {filteredEmployees.length > 0 ? (
              filteredEmployees.map((emp) => {
                const isEditing = editingId === emp.employeeNumber

                return (
                  <tr
                    key={emp.employeeNumber}
                    className={`hover:bg-slate-50/50 transition-colors ${
                      isEditing ? 'bg-primary/5 hover:bg-primary/5' : ''
                    }`}
                  >
                    {/* Column 1: Name & ID */}
                    <td className="px-6 py-4.5">
                      <div className="flex items-center gap-3">
                        {/* Avatar */}
                        <div
                          className={`w-9 h-9 rounded-full flex items-center justify-center font-bold text-xs shrink-0 select-none ${emp.avatarColor}`}
                        >
                          {getInitials(emp.firstName, emp.lastName)}
                        </div>

                        {/* Name Info */}
                        <div className="flex flex-col gap-1 w-full">
                          {isEditing ? (
                            <div className="flex gap-1.5 w-full">
                              <input
                                type="text"
                                name="firstName"
                                value={editFormData.firstName}
                                onChange={handleEditChange}
                                placeholder="Nombre"
                                className={
                                  validationErrors.firstName ? errorInputClass : inputClass
                                }
                              />
                              <input
                                type="text"
                                name="lastName"
                                value={editFormData.lastName}
                                onChange={handleEditChange}
                                placeholder="Apellidos"
                                className={validationErrors.lastName ? errorInputClass : inputClass}
                              />
                            </div>
                          ) : (
                            <span className="font-semibold text-neutral text-sm leading-tight">
                              {emp.firstName} {emp.lastName}
                            </span>
                          )}
                          <span className="text-[10px] text-secondary font-medium tracking-wide">
                            ID: EMP-{emp.employeeNumber}
                          </span>
                        </div>
                      </div>
                    </td>

                    {/* Column 2: Job & Department */}
                    <td className="px-6 py-4.5">
                      {isEditing ? (
                        <div className="flex flex-col gap-1.5 max-w-[180px]">
                          <select
                            name="jobTitle"
                            value={editFormData.jobTitle}
                            onChange={handleEditChange}
                            className={validationErrors.jobTitle ? errorInputClass : inputClass}
                          >
                            <option value="" disabled>
                              Seleccionar Rol
                            </option>
                            {roles.map((r) => (
                              <option key={r.value} value={r.value}>
                                {r.label}
                              </option>
                            ))}
                          </select>
                          <select
                            name="department"
                            value={editFormData.department}
                            onChange={handleEditChange}
                            className={validationErrors.department ? errorInputClass : inputClass}
                          >
                            <option value="" disabled>
                              Seleccionar Departamento
                            </option>
                            {departments.map((d) => (
                              <option key={d.value} value={d.value}>
                                {d.label}
                              </option>
                            ))}
                          </select>
                        </div>
                      ) : (
                        <div className="flex flex-col leading-tight">
                          <span className="font-medium text-neutral">
                            {roles.find((r) => r.value === emp.jobTitle)?.label || emp.jobTitle}
                          </span>
                          <span className="text-xs text-secondary mt-0.5">
                            {departments.find((d) => d.value === emp.department)?.label ||
                              emp.department}
                          </span>
                        </div>
                      )}
                    </td>

                    {/* Column 3: Contact */}
                    <td className="px-6 py-4.5">
                      {isEditing ? (
                        <input
                          type="email"
                          name="email"
                          value={editFormData.email}
                          onChange={handleEditChange}
                          placeholder="e.rodriguez@corp.com"
                          className={
                            validationErrors.email
                              ? errorInputClass
                              : `${inputClass} w-full max-w-[200px]`
                          }
                        />
                      ) : (
                        <div className="flex items-center gap-1.5 text-secondary hover:text-neutral transition-colors">
                          <Mail size={13} className="shrink-0 text-slate-400" />
                          <span className="text-xs font-medium">{emp.email}</span>
                        </div>
                      )}
                    </td>

                    {/* Column 4: Actions */}
                    <td className="px-6 py-4.5 text-right">
                      {isEditing ? (
                        <div className="flex items-center justify-end gap-1.5">
                          <button
                            onClick={() => handleEditSave(emp.employeeNumber)}
                            title="Guardar cambios"
                            className="p-1 text-emerald-600 hover:bg-emerald-50 rounded transition-colors cursor-pointer"
                          >
                            <Check size={16} />
                          </button>
                          <button
                            onClick={handleEditCancel}
                            title="Cancelar edición"
                            className="p-1 text-red-500 hover:bg-red-50 rounded transition-colors cursor-pointer"
                          >
                            <X size={16} />
                          </button>
                        </div>
                      ) : (
                        <button
                          onClick={() => handleEditStart(emp)}
                          disabled={editingId !== null}
                          title={
                            editingId !== null
                              ? 'Guarde o cancele la edición actual'
                              : 'Editar empleado'
                          }
                          className={`p-1.5 rounded transition-all inline-flex items-center justify-center ${
                            editingId !== null
                              ? 'text-slate-300 cursor-not-allowed'
                              : 'text-secondary hover:text-neutral hover:bg-slate-100 cursor-pointer'
                          }`}
                        >
                          <Pencil size={14} />
                        </button>
                      )}
                    </td>
                  </tr>
                )
              })
            ) : (
              <tr>
                <td colSpan="5" className="px-6 py-12 text-center text-secondary">
                  No se encontraron empleados que coincidan con la búsqueda.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination Footer */}
      <div className="p-4 border-t border-slate-200 bg-slate-50 flex items-center justify-between text-xs text-secondary font-medium select-none">
        <div>
          Mostrando 1 a {filteredEmployees.length} de {employees.length} empleados activos
        </div>
        <div className="flex items-center gap-1">
          <button
            disabled
            className="p-1.5 border border-slate-200 bg-white text-slate-300 rounded hover:bg-slate-50 disabled:opacity-50 disabled:hover:bg-white cursor-not-allowed"
          >
            <ChevronLeft size={14} />
          </button>
          <button
            disabled
            className="p-1.5 border border-slate-200 bg-white text-slate-300 rounded hover:bg-slate-50 disabled:opacity-50 disabled:hover:bg-white cursor-not-allowed"
          >
            <ChevronRight size={14} />
          </button>
        </div>
      </div>
    </div>
  )
}
