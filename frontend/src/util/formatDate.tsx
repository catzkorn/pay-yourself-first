function formatDateAsDay(dateString: string): string {
  const date = new Date(dateString);
  let d = date.getMonth() + 1 + "/" + date.getFullYear();
  return d;
}

export default formatDateAsDay;
