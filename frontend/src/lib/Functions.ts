// format date to YYYY-MM-DD (for date input)
export const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
};

// format flaot to 2 decimal places
export const formatFloat = (float: number, decimalBit = 2) => {
  if (decimalBit < 0) {
    return float;
  }
  if (decimalBit === 0) {
    return Math.round(float);
  }
  return float.toFixed(decimalBit);
};
