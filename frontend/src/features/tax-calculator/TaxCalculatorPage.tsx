import { useState } from 'react'
import type { FormEvent } from 'react'
import { SectionTitle } from '../../components/SectionTitle'
import { calculateTax } from '../../services/taxApi'

export function TaxCalculatorPage() {
  const [annualIncome, setAnnualIncome] = useState('')
  const [taxOwed, setTaxOwed] = useState<number | null>(null)
  const [effectiveTaxRate, setEffectiveTaxRate] = useState<number | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const parsedIncome = Number(annualIncome)
    if (Number.isNaN(parsedIncome) || parsedIncome < 0) {
      setError('Please enter a valid annual income.')
      setTaxOwed(null)
      setEffectiveTaxRate(null)
      return
    }

    setLoading(true)
    setError(null)

    try {
      const result = await calculateTax({
        annualIncome: parsedIncome,
        taxYear: '2025-2026',
        countryCode: 'AU',
      })
      setTaxOwed(result.taxOwed)
      setEffectiveTaxRate(result.effectiveTaxRate)
    } catch (submissionError) {
      setTaxOwed(null)
      setEffectiveTaxRate(null)
      setError(submissionError instanceof Error ? submissionError.message : 'Failed to calculate tax.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <section className="tax-calculator">
      <SectionTitle>Tax Calculator</SectionTitle>
      <form className="tax-calculator-form" onSubmit={onSubmit}>
        <label htmlFor="annualIncome">Annual income (AUD)</label>
        <input
          id="annualIncome"
          type="number"
          min="0"
          step="100"
          value={annualIncome}
          onChange={(event) => setAnnualIncome(event.target.value)}
          required
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Calculating...' : 'Calculate'}
        </button>
      </form>
      {taxOwed !== null && effectiveTaxRate !== null && (
        <p className="tax-result">
          Estimated tax owed: ${taxOwed.toFixed(2)} ({effectiveTaxRate.toFixed(2)}%)
        </p>
      )}
      {error && <p className="tax-error">{error}</p>}
    </section>
  )
}
