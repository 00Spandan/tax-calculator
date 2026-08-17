import { describe, expect, it, vi } from 'vitest'
import { calculateTax } from './taxApi'

describe('calculateTax', () => {
  it('returns parsed tax response', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ taxableIncome: 50000, taxOwed: 6717, effectiveTaxRate: 13.434 }),
      }),
    )

    const result = await calculateTax({ annualIncome: 50000 })

    expect(result.taxOwed).toBe(6717)
  })
})
