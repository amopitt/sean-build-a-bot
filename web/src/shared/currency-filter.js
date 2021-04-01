export default function currencyFilter(amount, symbol) {
  return `${symbol}${amount.toFixed(2)}`;
}
