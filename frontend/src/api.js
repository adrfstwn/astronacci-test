const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

async function post(path, body) {
  const res = await fetch(`${BASE_URL}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })

  const data = await res.json().catch(() => ({}))

  if (!res.ok) {
    throw new Error(data.error || `Request failed with status ${res.status}`)
  }

  return data
}

export function checkVoucher({ flightNumber, date }) {
  return post('/api/check', { flightNumber, date })
}

export function generateVoucher({ name, id, flightNumber, date, aircraft }) {
  return post('/api/generate', { name, id, flightNumber, date, aircraft })
}
