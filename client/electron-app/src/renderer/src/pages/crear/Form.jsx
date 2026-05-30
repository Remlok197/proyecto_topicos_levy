import React, { useState } from 'react'
import {
  User,
  Briefcase,
  Save,
  Calendar,
  ChevronDown,
  AlertCircle,
  CheckCircle2,
  RotateCcw
} from 'lucide-react'

export default function Form() {
  // Form state as requested
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    birthDate: '',
    department: '',
    role: '',
    initialSalary: ''
  })

  // Error state for validation
  const [errors, setErrors] = useState({})
  // Submission success state for mockup presentation
  const [successData, setSuccessData] = useState(null)

  // Field change handler
  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }))

    // Clear validation error when user types
    if (errors[name]) {
      setErrors((prev) => ({
        ...prev,
        [name]: null
      }))
    }
  }

  // Form validation
  const validateForm = () => {
    const newErrors = {}

    if (!formData.firstName.trim()) {
      newErrors.firstName = 'Campo requerido'
    }
    if (!formData.lastName.trim()) {
      newErrors.lastName = 'Campo requerido'
    }
    if (!formData.birthDate) {
      newErrors.birthDate = 'Campo requerido'
    }
    if (!formData.department) {
      newErrors.department = 'Campo requerido'
    }
    if (!formData.role) {
      newErrors.role = 'Campo requerido'
    }
    if (!formData.initialSalary.trim()) {
      newErrors.initialSalary = 'Campo requerido'
    } else if (isNaN(formData.initialSalary) || Number(formData.initialSalary) <= 0) {
      newErrors.initialSalary = 'El salario debe ser un número válido mayor a 0'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  // Form submission handler
  const handleSubmit = (e) => {
    e.preventDefault()

    if (validateForm()) {
      // Form is valid! Set success data for front-end demonstration
      setSuccessData(formData)
    }
  }

  // Cancel handler - resets the form and errors
  const handleCancel = () => {
    setFormData({
      firstName: '',
      lastName: '',
      birthDate: '',
      department: '',
      role: '',
      initialSalary: ''
    })
    setErrors({})
    setSuccessData(null)
  }

  // Common classes
  const labelClass = 'text-xs font-semibold text-neutral/80 block mb-1.5 transition-colors'
  const inputBaseClass =
    'w-full bg-slate-50 border text-neutral placeholder-slate-400 rounded-lg py-2.5 px-3.5 text-sm transition-all focus:bg-white focus:outline-none focus:ring-2'
  const inputNormalClass = `${inputBaseClass} border-slate-200 focus:ring-primary/20 focus:border-primary`
  const inputErrorClass = `${inputBaseClass} border-red-500 focus:ring-red-100 focus:border-red-500`

  // Departments list (English values, Spanish labels)
  const departments = [
    { value: 'HR', label: 'Recursos Humanos (HR)' },
    { value: 'IT', label: 'Tecnología (IT)' },
    { value: 'Sales', label: 'Ventas (Sales)' },
    { value: 'Finance', label: 'Finanzas (Finance)' },
    { value: 'Marketing', label: 'Mercadotecnia (Marketing)' }
  ]

  // Roles list
  const roles = [
    { value: 'Developer', label: 'Desarrollador (Developer)' },
    { value: 'Designer', label: 'Diseñador (Designer)' },
    { value: 'Manager', label: 'Gerente (Manager)' },
    { value: 'Analyst', label: 'Analista (Analyst)' },
    { value: 'Specialist', label: 'Especialista (Specialist)' }
  ]

  // If form was successfully submitted, display details
  if (successData) {
    return (
      <div className="bg-white border border-slate-200 rounded-xl p-8 shadow-sm max-w-4xl w-full flex flex-col items-center justify-center text-center animate-fade-in">
        <div className="bg-emerald-50 text-emerald-600 p-4 rounded-full mb-4">
          <CheckCircle2 size={48} />
        </div>
        <h2 className="text-xl font-bold text-neutral mb-2">¡Empleado Registrado con Éxito!</h2>
        <p className="text-secondary text-sm mb-6 max-w-md">
          Los datos del empleado han sido validados localmente y están listos para enviarse a la API
          de backend.
        </p>

        <div className="w-full max-w-lg bg-slate-50 border border-slate-200 rounded-lg p-5 text-left mb-6">
          <h3 className="font-semibold text-neutral border-b border-slate-200 pb-2 mb-3">
            Resumen de Datos
          </h3>
          <div className="grid grid-cols-2 gap-y-3 gap-x-4 text-sm">
            <div>
              <span className="text-secondary block text-xs">Nombre</span>
              <span className="text-neutral font-medium">{successData.firstName}</span>
            </div>
            <div>
              <span className="text-secondary block text-xs">Apellidos</span>
              <span className="text-neutral font-medium">{successData.lastName}</span>
            </div>
            <div>
              <span className="text-secondary block text-xs">Fecha de Nacimiento</span>
              <span className="text-neutral font-medium">{successData.birthDate}</span>
            </div>
            <div>
              <span className="text-secondary block text-xs">Salario Inicial (Anual)</span>
              <span className="text-neutral font-medium">
                ${Number(successData.initialSalary).toLocaleString()}
              </span>
            </div>
            <div>
              <span className="text-secondary block text-xs">Departamento</span>
              <span className="text-neutral font-medium">
                {departments.find((d) => d.value === successData.department)?.label ||
                  successData.department}
              </span>
            </div>
            <div>
              <span className="text-secondary block text-xs">Rol</span>
              <span className="text-neutral font-medium">
                {roles.find((r) => r.value === successData.role)?.label || successData.role}
              </span>
            </div>
          </div>
        </div>

        <button
          onClick={handleCancel}
          className="flex items-center gap-2 bg-primary hover:bg-primary/95 text-white py-2.5 px-6 rounded-lg font-medium text-sm transition-colors shadow-sm cursor-pointer"
        >
          <RotateCcw size={16} />
          Registrar Otro Empleado
        </button>
      </div>
    )
  }

  return (
    <div className="bg-white border border-slate-200 rounded-xl shadow-sm max-w-4xl w-full overflow-hidden">
      <form onSubmit={handleSubmit} className="p-8 flex flex-col gap-8">
        {/* SECTION 1: Personal Information */}
        <div>
          <div className="flex items-center gap-2.5 text-primary">
            <User size={20} className="stroke-[2.5]" />
            <h2 className="text-base font-bold tracking-tight">Información Personal</h2>
          </div>
          <hr className="border-slate-100 mt-2.5 mb-5" />

          <div className="flex flex-col gap-5">
            {/* Grid for Name & Last Name */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-5">
              {/* First Name */}
              <div>
                <label htmlFor="firstName" className={labelClass}>
                  Nombre
                </label>
                <input
                  type="text"
                  id="firstName"
                  name="firstName"
                  placeholder="ej. María"
                  value={formData.firstName}
                  onChange={handleChange}
                  className={errors.firstName ? inputErrorClass : inputNormalClass}
                />
                {errors.firstName && (
                  <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                    <AlertCircle size={12} className="fill-red-600 text-white" />
                    {errors.firstName}
                  </span>
                )}
              </div>

              {/* Last Name */}
              <div>
                <label htmlFor="lastName" className={labelClass}>
                  Apellidos
                </label>
                <input
                  type="text"
                  id="lastName"
                  name="lastName"
                  placeholder="ej. García"
                  value={formData.lastName}
                  onChange={handleChange}
                  className={errors.lastName ? inputErrorClass : inputNormalClass}
                />
                {errors.lastName && (
                  <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                    <AlertCircle size={12} className="fill-red-600 text-white" />
                    {errors.lastName}
                  </span>
                )}
              </div>
            </div>

            {/* Birth Date */}
            <div>
              <label htmlFor="birthDate" className={labelClass}>
                Fecha de Nacimiento
              </label>
              <div className="relative">
                <input
                  type="date"
                  id="birthDate"
                  name="birthDate"
                  value={formData.birthDate}
                  onChange={handleChange}
                  className={`${errors.birthDate ? inputErrorClass : inputNormalClass} pr-10`}
                />
                {!formData.birthDate && (
                  <Calendar
                    size={16}
                    className="absolute right-3.5 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none"
                  />
                )}
              </div>
              {errors.birthDate && (
                <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                  <AlertCircle size={12} className="fill-red-600 text-white" />
                  {errors.birthDate}
                </span>
              )}
            </div>
          </div>
        </div>

        {/* SECTION 2: Employment Details */}
        <div>
          <div className="flex items-center gap-2.5 text-primary">
            <Briefcase size={20} className="stroke-[2.5]" />
            <h2 className="text-base font-bold tracking-tight">Detalles de Empleo</h2>
          </div>
          <hr className="border-slate-100 mt-2.5 mb-5" />

          <div className="flex flex-col gap-5">
            {/* Grid for Department & Role */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-5">
              {/* Department */}
              <div>
                <label htmlFor="department" className={labelClass}>
                  Departamento
                </label>
                <div className="relative">
                  <select
                    id="department"
                    name="department"
                    value={formData.department}
                    onChange={handleChange}
                    className={`${errors.department ? inputErrorClass : inputNormalClass} appearance-none pr-10`}
                  >
                    <option value="" disabled>
                      Seleccionar Departamento
                    </option>
                    {departments.map((dept) => (
                      <option key={dept.value} value={dept.value}>
                        {dept.label}
                      </option>
                    ))}
                  </select>
                  <ChevronDown
                    className="absolute right-3.5 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none"
                    size={16}
                  />
                </div>
                {errors.department && (
                  <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                    <AlertCircle size={12} className="fill-red-600 text-white" />
                    {errors.department}
                  </span>
                )}
              </div>

              {/* Role */}
              <div>
                <label htmlFor="role" className={labelClass}>
                  Rol
                </label>
                <div className="relative">
                  <select
                    id="role"
                    name="role"
                    value={formData.role}
                    onChange={handleChange}
                    className={`${errors.role ? inputErrorClass : inputNormalClass} appearance-none pr-10`}
                  >
                    <option value="" disabled>
                      Seleccionar Rol
                    </option>
                    {roles.map((role) => (
                      <option key={role.value} value={role.value}>
                        {role.label}
                      </option>
                    ))}
                  </select>
                  <ChevronDown
                    className="absolute right-3.5 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none"
                    size={16}
                  />
                </div>
                {errors.role && (
                  <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                    <AlertCircle size={12} className="fill-red-600 text-white" />
                    {errors.role}
                  </span>
                )}
              </div>
            </div>

            {/* Salary */}
            <div>
              <label htmlFor="initialSalary" className={labelClass}>
                Salario Inicial (Anual)
              </label>
              <div className="relative">
                <span className="absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400 text-sm select-none">
                  $
                </span>
                <input
                  type="number"
                  id="initialSalary"
                  name="initialSalary"
                  placeholder="65000"
                  value={formData.initialSalary}
                  onChange={handleChange}
                  className={`${errors.initialSalary ? inputErrorClass : inputNormalClass} pl-8`}
                />
              </div>
              {errors.initialSalary && (
                <span className="flex items-center gap-1 text-[11px] font-medium text-red-600 mt-1.5">
                  <AlertCircle size={12} className="fill-red-600 text-white" />
                  {errors.initialSalary}
                </span>
              )}
            </div>
          </div>
        </div>

        {/* Divider above buttons */}
        <hr className="border-slate-100 mt-2" />

        {/* Action Buttons */}
        <div className="flex items-center justify-end gap-3 -mt-3">
          <button
            type="button"
            onClick={handleCancel}
            className="border border-slate-300 text-neutral hover:bg-slate-50 transition-colors py-2.5 px-6 rounded-lg font-medium text-sm cursor-pointer"
          >
            Cancelar
          </button>
          <button
            type="submit"
            className="bg-primary hover:bg-primary/95 text-white flex items-center gap-2 py-2.5 px-6 rounded-lg font-medium text-sm transition-colors shadow-sm cursor-pointer"
          >
            <Save size={16} />
            Guardar
          </button>
        </div>
      </form>
    </div>
  )
}
