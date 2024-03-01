export default function formatCurrency(price: number) {
  return price.toLocaleString("pt-br", {
    style: "currency",
    currency: "BRL",
  });
}
