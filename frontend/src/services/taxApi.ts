export interface CalculateTaxRequest {
  annualIncome: number
  taxYear?: string
  countryCode?: string
}

export interface CalculateTaxResponse {
  taxableIncome: number
  taxOwed: number
  effectiveTaxRate: number
}

interface ErrorResponse {
  error?: string
}

export async function calculateTax(payload: CalculateTaxRequest): Promise<CalculateTaxResponse> {
  const response = await fetch('/api/tax/calculate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const errorBody = (await response.json().catch(() => ({}))) as ErrorResponse
    throw new Error(errorBody.error ?? 'Tax calculation request failed.')
  }

  return (await response.json()) as CalculateTaxResponse
}
