/**
 * Calculate modulo. Javascript's `%` operator is a remainder operator.
 *
 * @param n - dividend
 * @param d - divisor
 * @returns `n % d` but if `%` was a modulo operator:
 */
export const mod = (n: number, d: number) => ((n % d) + d) % d;
