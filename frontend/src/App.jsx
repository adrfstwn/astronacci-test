import { useState } from 'react'
import { checkVoucher, generateVoucher } from './api'

const AIRCRAFT_OPTIONS = ['ATR', 'Airbus 320', 'Boeing 737 Max']

const initialForm = {
  crewName: '',
  crewId: '',
  flightNumber: '',
  flightDate: '',
  aircraft: AIRCRAFT_OPTIONS[0],
}

function toDDMMYYYY(isoDate) {
  const [year, month, day] = isoDate.split('-')
  return `${day}-${month}-${year}`
}

const inputClass =
  'w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 shadow-sm ' +
  'placeholder:text-slate-400 focus:border-violet-500 focus:outline-none focus:ring-2 focus:ring-violet-200 ' +
  'dark:border-slate-700 dark:bg-slate-800 dark:text-slate-100 dark:focus:border-violet-400 dark:focus:ring-violet-900'

function Field({ label, children }) {
  return (
    <label className="block space-y-1.5 text-sm font-medium text-slate-700 dark:text-slate-300">
      {label}
      {children}
    </label>
  )
}

function App() {
  const [form, setForm] = useState(initialForm)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [seats, setSeats] = useState(null)

  const handleChange = (e) => {
    const { name, value } = e.target
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    setSeats(null)
    setLoading(true)

    const date = toDDMMYYYY(form.flightDate)

    try {
      const { exists } = await checkVoucher({ flightNumber: form.flightNumber, date })

      if (exists) {
        setError(`Vouchers have already been generated for flight ${form.flightNumber} on ${date}.`)
        return
      }

      const result = await generateVoucher({
        name: form.crewName,
        id: form.crewId,
        flightNumber: form.flightNumber,
        date,
        aircraft: form.aircraft,
      })

      setSeats(result.seats)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex min-h-svh items-center justify-center bg-slate-50 px-4 py-10 dark:bg-slate-950">
      <div className="w-full max-w-md rounded-2xl border border-slate-200 bg-white p-8 shadow-xl shadow-slate-200/60 dark:border-slate-800 dark:bg-slate-900 dark:shadow-none">
        <div className="mb-6 text-center">
          <h1 className="text-xl font-semibold text-slate-900 dark:text-white">Voucher Generator</h1>
          <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
            Fill in the flight details to generate vouchers.
          </p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          <Field label="Crew Name">
            <input
              name="crewName"
              value={form.crewName}
              onChange={handleChange}
              className={inputClass}
              required
            />
          </Field>

          <Field label="Crew ID">
            <input name="crewId" value={form.crewId} onChange={handleChange} className={inputClass} required />
          </Field>

          <Field label="Flight Number">
            <input
              name="flightNumber"
              value={form.flightNumber}
              onChange={handleChange}
              placeholder="e.g. GA102"
              className={inputClass}
              required
            />
          </Field>

          <Field label="Flight Date">
            <input
              type="date"
              name="flightDate"
              value={form.flightDate}
              onChange={handleChange}
              className={inputClass}
              required
            />
          </Field>

          <Field label="Aircraft Type">
            <select name="aircraft" value={form.aircraft} onChange={handleChange} className={inputClass} required>
              {AIRCRAFT_OPTIONS.map((option) => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </select>
          </Field>

          <button
            type="submit"
            disabled={loading}
            className="w-full rounded-lg bg-violet-600 px-4 py-2.5 text-sm font-semibold text-white shadow-sm transition hover:bg-violet-500 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {loading ? 'Generating…' : 'Generate Vouchers'}
          </button>
        </form>

        {error && (
          <div className="mt-5 rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700 dark:border-red-900/50 dark:bg-red-950/40 dark:text-red-300">
            {error}
          </div>
        )}

        {seats && (
          <div className="mt-5 rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-800 dark:border-emerald-900/50 dark:bg-emerald-950/40 dark:text-emerald-300">
            <p>Vouchers generated successfully. Assigned seats:</p>
            <ul className="mt-2 flex flex-wrap gap-2">
              {seats.map((seat) => (
                <li
                  key={seat}
                  className="rounded-md bg-emerald-100 px-2.5 py-1 font-mono text-xs font-semibold text-emerald-900 dark:bg-emerald-900/60 dark:text-emerald-200"
                >
                  {seat}
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </div>
  )
}

export default App
